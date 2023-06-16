package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

// fileSizes
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

var verbose = flag.Bool(
	"v",
	false,
	"show verbose progress messages",
)

func main() {
	start := time.Now()
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64) // bidirectional channel
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}

		close(fileSizes)
	}()

	var tick <-chan time.Time // receive-only channel

	/*
		<- chan -- receive only. reference point -- you
		chan <- -- send only. reference point -- you
	*/

	if *verbose {
		/*
			time.Tick creates a new channel that
			sends the current time everyspecified interval.
		*/
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles += 1
			nbytes += size
		case <-tick: // receiving time, value not important
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
	fmt.Println(time.Now().UnixMilli() - start.UnixMilli())
}
