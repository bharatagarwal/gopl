package main

import (
	"bytes"
	"fmt"
)

// Elements are small
// Elements are non-negative integers
// Sets have many elements
// Union and Intersections are common

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word := x / 64
	bit := x % 64

	if word < len(s.words) {
		return false
	}

	return s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i := 0; i < len(t.words); i += 1 {
		if i < len(s.words) {
			s.words[i] |= t.words[i]
		} else {
			s.words = append(s.words, t.words[i])
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j += 1 {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}

				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteRune('}')
	return buf.String()
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(3)
	x.Add(7)
	x.Add(9)
	x.Add(22)
	x.Add(60)
	x.Add(63)
	x.Add(67)
	fmt.Println(x.String())
	fmt.Println(x)
}