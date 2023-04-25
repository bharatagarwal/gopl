package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var counts = make(map[rune]int)
var utfLen [utf8.UTFMax + 1]int
var invalid int

func parseContent(in *bufio.Reader) {
	char, size, err := in.ReadRune()

	for err != io.EOF {
		char, size, err = in.ReadRune()

		if err != nil {
			_, err2 := fmt.Fprintf(os.Stderr, "charcount: %v\n", err)

			if err2 != nil {
				fmt.Println("Error printing error code", err2)
			}

			os.Exit(1)
		}

		if char == unicode.ReplacementChar && size == 1 {
			invalid += 1
			continue
		}

		counts[char] += 1
		utfLen[size] += 1
	}
}

func printRuneCount() {
	fmt.Print("rune\tcount\n")

	for char, count := range counts {
		fmt.Printf("%q\t%d\n", char, count)
	}
}

func printLengthCount() {
	fmt.Print("\nlen\tcount\n")

	for length, count := range utfLen {
		if length > 0 {
			fmt.Printf("%d\t%d\n", length, count)
		}
	}
}

func printInvalidCount() {
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters", invalid)
	}
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	parseContent(stdin)
	printRuneCount()
	printLengthCount()
	printInvalidCount()
}