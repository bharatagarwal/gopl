package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire
	defer func() { <-sema }() // release

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

func walkDir(
	dir string,
	wg *sync.WaitGroup,
	fileSizes chan<- int64,
) {
	defer wg.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, wg, fileSizes)
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

var vFlag *bool = flag.Bool(
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

	fileSizes := make(chan int64)
	var wg sync.WaitGroup

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	var nfiles, nbytes int64

	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
	fmt.Println(time.Now().UnixMilli() - start.UnixMilli())
}