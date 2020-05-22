package main

import (
	"fmt"
	"time"
)

//ch关闭后 for...range读取剩余数据后退出
func main() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	for value := range ch {
		fmt.Printf("value: %d\n", value)
	}
}
