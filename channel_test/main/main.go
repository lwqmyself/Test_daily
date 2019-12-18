package main

import "fmt"

var res chan string = make(chan string, 20)
var boolChanPro chan bool = make(chan bool, 2)
var boolChanCos chan bool = make(chan bool, 2)

func Producer(n int) {
	for i := 1; i <= 10; i++ {
		res <- fmt.Sprintf("生产者 %d 生产的 %d", n, i)
	}
	boolChanPro <- true
}

func Cosumer(n int) {
	for {
		v, ok := <-res
		if ok {
			fmt.Printf("消费者%d消费了%s\n", n, v)
		} else {
			break
		}
	}
	boolChanCos <- true
}

func main() {
	for i := 1; i <= 2; i++ {
		go Producer(i)
	}
	for i := 1; i <= 2; i++ {
		go Cosumer(i)
	}

	go func() {
		for i := 1; i <= 2; i++ {
			<-boolChanPro
		}
		close(res)
	}()

	for i := 1; i <= 2; i++ {
		<-boolChanCos
	}
}
