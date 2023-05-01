package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

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

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case bailout{}: // expected panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(node *html.Node) {
		if node.Type == html.ElementNode &&
			node.Data == "title" &&
			node.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}

			title = node.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}

func main() {
	url := "https://example.com" // replace with the desired URL

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML from %s: %v", url, err)
	}

	title, err := soleTitle(doc)
	if err != nil {
		log.Fatalf("Error getting title: %v", err)
	}

	fmt.Println("Title:", title)
}