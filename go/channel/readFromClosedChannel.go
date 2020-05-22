package main

import (
	"fmt"
	"time"
)

//数据上游关闭ch后，下游可以监听ok状态，处理完成后退出，实现安全退出//ch关闭后读取到的为类型0值和false
func main() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	for {
		value, ok := <-ch
		fmt.Printf("value: %d,ok: %v\n", value, ok)
		if !ok {
			return
		}
		time.Sleep(time.Second)
	}
}
