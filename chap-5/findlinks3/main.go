package main

import (
	"fmt"
	"log"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		fmt.Printf("%#v\n", worklist)
		items := worklist
		worklist = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
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
	link := []string{"https://go.dev"}
	breadthFirst(crawl, link)
}