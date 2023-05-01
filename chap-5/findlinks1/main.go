package main

import (
	"fmt"
	"log"
	"os"
)

import (
	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attribute := range n.Attr {
			if attribute.Key == "href" {
				links = append(links, attribute.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}