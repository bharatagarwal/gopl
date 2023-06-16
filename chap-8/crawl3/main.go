package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := getHttpResponse(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := parseHtml(resp.Body)
	if err != nil {
		return nil, err
	}

	links, err := extractLinks(doc, resp.Request.URL)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	// Traverse each child of the current node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func getHttpResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	return resp, nil
}

func parseHtml(body io.Reader) (*html.Node, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil,
			fmt.Errorf("parsing HTML: %v", err)
	}

	return doc, nil
}

func extractLinks(doc *html.Node, base *url.URL) ([]string, error) {
	var links []string

	// applied to each HTML node by forEachNode function
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := base.Parse(a.Val)
				if err != nil {
					continue
				}

				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i += 1 {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)

	for item := range worklist {
		for _, link := range item {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link // sends deduplicated, unseen links
			}
		}
	}
}