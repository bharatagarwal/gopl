package main

import (
	"fmt"
	"sort"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(courses []string)

	visitAll = func(courses []string) {
		for _, course := range courses {
			if !seen[course] {
				seen[course] = true
				visitAll(m[course])
				order = append(order, course)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}

func main() {
	sortedCourses := topoSort(prereqs)
	fmt.Println(strings.Join(sortedCourses, "\n"))
}