package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 100)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	fmt.Printf("len(ch): %d,cap(ch): %d\n", len(ch), cap(ch))
	<-ch
	<-ch
	fmt.Printf("len(ch): %d,cap(ch): %d\n", len(ch), cap(ch))
}
