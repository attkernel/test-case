package main

import (
	"fmt"
)

//panic
func main() {
	var ch chan int
	go func() {
		ch <- 1
	}()
	fmt.Printf("value: %d\n", <-ch)
}
