package main

import (
	"fmt"
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

func anchorNode(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf(
				"getting %s: %s; error closing response body %s", url,
				resp.Status, err,
			)
		}

		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("getting %s: %s", url, err)
	}

	var links []string

	visitNode := func(n *html.Node) {
		if anchorNode(n) {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(attr.Val)

				if err != nil {
					continue // ignore bad URLs
				}

				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}