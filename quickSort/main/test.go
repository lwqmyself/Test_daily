package main

import (
	"fmt"
)

var numChan = make(chan int, 2000)
var resChan = make(chan int, 2000)
var boolChan = make(chan bool, 1)

func test() {

	go insert()
	for i := 1; i <= 8; i++ {
		go sum()
	}
	if <-boolChan {
		for v := range resChan {
			fmt.Println(v)
		}
	}

}

func insert() {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func sum() {
	for {
		v, ok := <-numChan
		res := 0
		for i := 1; i <= v; i++ {
			res += i
		}
		resChan <- res
		if !ok {
			break
		}
	}
	close(resChan)
	boolChan <- true
	close(boolChan)
}
