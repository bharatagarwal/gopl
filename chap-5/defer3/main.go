package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	if r := recover(); r != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		os.Stdout.Write(buf[:n])
		panic(r) // Re-throw the panic to get the default stack trace
	}
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}