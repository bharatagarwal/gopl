package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth += 1
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth -= 1
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}

func main() {
	url := "https://gopl.io"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "fetch: getting %s: %s\n", url, resp.Status)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}