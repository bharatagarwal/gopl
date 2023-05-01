package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, written int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)

	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, nil
	}

	written, err = io.Copy(f, resp.Body)

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, written, err
}

func main() {
	url := "https://go.dev/doc/"

	filename, written, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching %s: %v\n", url, err)
	}

	fmt.Printf("Downloaded %s to %s (%d bytes)\n", url, filename, written)
}