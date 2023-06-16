package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

func main() {
	fmt.Println("Commencing countdown")

	tick := time.Tick(time.Second)

	for i := 10; i > 0; i -= 1 {
		fmt.Println(i)
		<-tick // one second between each receive?
		launch()
	}
}