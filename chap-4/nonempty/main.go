package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i += 1
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Println(data)
	fmt.Println(nonempty(data))
	fmt.Println(data)
}