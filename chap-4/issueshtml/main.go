package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const IssuesURL = "https://api.github.com/search/issues"

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()

		if err != nil {
			return nil, fmt.Errorf("search query failed: %s and body failed to close", resp.Status)
		}
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err2 := resp.Body.Close()

		if err2 != nil {
			return nil, err2
		}

		return nil, err
	}

	err = resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	handle(err)

	err = issueList.Execute(os.Stdout, result)
	handle(err)
}