package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("B", "D")

	fmt.Println(hasEdge("B", "D"))
	fmt.Println(hasEdge("D", "B"))

	for k, v := range graph {
		fmt.Println(k, v)
	}
}