package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
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

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)

	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s at HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)

		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			if err != nil {
				log.Fatal(err)
			}

			continue
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}
}