package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(5)
	fmt.Println(runtime.NumCPU())
	ch1 := make(chan int, 1)
	ch1 <- 10
	close(ch1)
	<-ch1
	<-ch1
	x := 190
	for x != 0 {
		y := x % 10
		fmt.Println(y)
		x /= 10
	}
}
