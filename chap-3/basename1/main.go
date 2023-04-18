package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i -= 1 {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i -= 1 {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	s := "bharat/agarwal/main.go"
	fmt.Println(basename(s))
}
