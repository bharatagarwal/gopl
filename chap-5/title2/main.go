package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")

	if ct != "text/html" && strings.HasPrefix(ct, "text/html;") == false {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		_ = fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode &&
			n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)
	return nil
}

func main() {
	url := "https://golang.org/doc/gopher/frontpage.png"
	err := title(url)
	if err != nil {
		log.Fatal(err)
	}

	os.ReadFile()
}