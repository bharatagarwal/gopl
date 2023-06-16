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

// calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n.
// Both functions are optional.
// pre is called before the children are visited (preorder)
// post is called after (postorder).
func forEachNode(
	n *html.Node,
	pre, post func(n *html.Node),
) {
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

// Extract makes a GET request, parses the response as HTML
// and returns the links in the HTML document.
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

// parses an HTML document and returns the root node
func parseHtml(body io.Reader) (*html.Node, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil,
			fmt.Errorf("parsing HTML: %v", err)
	}

	return doc, nil
}

func extractLinks(
	doc *html.Node,
	base *url.URL,
) ([]string, error) {
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

var tokens = make(chan struct{}, 20)

// Prints the URL, extracts
// the links and returns them.
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // send to channel -- will this panic,
	// or will this block?
	list, err := Extract(url)
	<-tokens // receive from channel
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)

	// counter -- number of pending sends to worklist
	var pending int

	pending += 1
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for ; pending > 0; pending -= 1 {
		item := <-worklist

		for _, link := range item {
			if !seen[link] {
				fmt.Printf("Pending: %d\n", pending)
				seen[link] = true
				pending += 1

				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}