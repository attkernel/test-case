package main

//panic
func main() {
	var ch chan int
	close(ch)
}
