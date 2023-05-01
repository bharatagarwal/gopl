package main

import "fmt"

func sum(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}

	return total
}

func f(...int) {}
func g([]int)  {}

func main() {
	fmt.Println(sum(1, 2, 3))
	values := []int{1, 2, 3}
	fmt.Println(sum(values...))

	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

}