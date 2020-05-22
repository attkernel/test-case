package main

import (
	"fmt"
	"time"
)

//close ch会唤醒所有等待在该 channel 上的 g，并使其进入 Grunnable 状态，这时这些 writer goroutine 会发现该 channel 已经是 closed 状态，就 panic了。
func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)
	close(ch)
	fmt.Printf("value: %d\n", <-ch)
}
