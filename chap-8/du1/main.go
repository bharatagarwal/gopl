package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
	FileInfo interface
	 - Name() string
	 - Size() int64
	 - Mode() FileMode
	 - ModTime() time.Time
	 - IsDir() bool

Reads a directory and returns a list of fs.FileInfo
sorted by filename.
*/
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

// sends filesize to channel, otherwise invokes itself if
// the file is a directory
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf(
		"%d files  %.1f GB\n",
		nfiles,
		float64(nbytes)/1e9,
	)
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)

	// walkDir would be a blocking operation if it was
	// invoked as part of the main goroutine
	go func() {
		// collects file sizes for given roots
		for _, root := range roots {
			walkDir(root, fileSizes)
		}

		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles += 1
		nbytes += size
	}

	printDiskUsage(nfiles, nbytes)
}