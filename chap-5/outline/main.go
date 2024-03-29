package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		if err != nil {
			log.Fatal(err)
		}
	}

	outline(nil, doc)
}