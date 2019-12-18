package main

import (
	"fmt"
)

var id int64

func main() {
	chanBegin := make(chan int, 1)
	go func() {
		id = id + 2
		fmt.Printf("in groutine id = %d\n", id)
	}()
	fmt.Printf("id = %d", id)
	v, ok := <-chanBegin
	if ok {
		fmt.Println(v)
	}
}
