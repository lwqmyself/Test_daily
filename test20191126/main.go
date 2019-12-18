package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
