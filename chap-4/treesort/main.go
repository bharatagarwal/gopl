package main

import "fmt"

var addInvocations = 0
var appendInvocations = 0

type tree struct {
	value       int
	left, right *tree
}

func appendValues(values []int, t *tree) []int {
	appendInvocations += 1
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func add(t *tree, value int) *tree {
	addInvocations += 1
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

func main() {
	sampleValues := []int{7, 3, 5, 8, 1, 10, 2, 6, 4, 9}
	fmt.Printf("Before sorting: %v\n", sampleValues)

	Sort(sampleValues)
	fmt.Printf("After sorting: %v\n", sampleValues)
	fmt.Println("Append", appendInvocations)
	fmt.Println("Add", addInvocations)
}