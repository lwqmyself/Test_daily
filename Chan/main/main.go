package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

var numChan = make(chan int, 200000)
var resChan = make(chan int, 201000)
var boolChan = make(chan bool, 80000)
var num1Chan = make(chan int, 2000)
var res1Chan = make(chan int, 2000)
var bool1Chan = make(chan bool, 80000)

func main() {
	//test()
	//Test1()
	/*var num string
	  fmt.Scanln(&num)
	  fmt.Println(strconv.Atoi(num[1:2]))*/
	test2()
}
func test() {
	timeNow := time.Now()
	go insert()
	for i := 1; i <= 80000; i++ {
		go sum()
	}
	//go sum()
	go func() {
		for i := 1; i <= 80000; i++ {
			<-boolChan
		}
		close(resChan)
		fmt.Println(len(resChan))
	}()

	time.Sleep(time.Second)
	ssss := make(map[int]int, 2000000)

	for {
		v, ok := <-resChan
		ssss[v] = v
		if !ok {
			break
		}
		//fmt.Println(v)
	}
	var keys []int
	for k := range ssss {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	i := -1
	for _, k := range keys {
		i++
		fmt.Println("Key:", i, "Value:", ssss[k])
	}
	t2 := time.Now()

	fmt.Println(t2.Sub(timeNow))
}

func insert() {
	for i := 1; i <= 200000; i++ {
		numChan <- i
		//fmt.Println(i)
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

		if !ok {

			break
		}
		if res > 0 {
			resChan <- res
		}
	}

	boolChan <- true

}

func Test1() {

	timeNow := time.Now()
	go shuixian()

	for i := 1; i <= 16; i++ {
		go sx2()
	}
	flag := make(chan bool)
	go func() {
		for i := 1; i <= 16; i++ {
			<-bool1Chan
		}
		close(res1Chan)
		fmt.Println("共", len(res1Chan), "个")
		flag <- true
	}()

	<-flag

	ssss := make(map[int]int, 50)

	for {
		v, ok := <-res1Chan
		ssss[v] = v
		if !ok {
			break
		}
		//fmt.Println(v)
	}
	var keys []int
	for k := range ssss {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	i := 0
	for _, k := range keys {
		if k == 0 {
			continue
		}
		i++
		fmt.Println(i, "----", ssss[k])
	}
	t2 := time.Now()

	fmt.Println(t2.Sub(timeNow))

}

func shuixian() {
	for i := 100; i <= 99999999; i++ {
		num1Chan <- i
		//fmt.Println(i)
	}
	close(num1Chan)
}
func sx2() {

	for {
		i, ok := <-num1Chan
		var len1 int = len(strconv.Itoa(i))
		var len2 int = len(strconv.Itoa(i))
		var s [20]int
		var sum = 0
		for j := 0; j < len1; j++ {
			s[j] = i / int(math.Pow10(len1-j-1)) % 10
			sum = sum + int(math.Pow(float64(s[j]), float64(len2)))
		}

		if !ok {

			break
		}
		if sum == i {
			res1Chan <- sum
		}
	}

	bool1Chan <- true
}

var num2Chan = make(chan int, 100000)
var res2Chan = make(chan int, 100000)
var flag2Chan = make(chan bool, 16)

func test2() {

	go inse()

	for i := 1; i <= 16; i++ {
		go su2Test()
	}
	flagb := make(chan bool, 1)
	go func() {
		for i := 1; i <= 16; i++ {
			<-flag2Chan
		}
		close(res2Chan)

		flagb <- true
		//time.Sleep(time.Second)

	}()

	//fmt.Println("共", len(res2Chan), "个")
	for {
		v, ok := <-res2Chan

		if !ok {
			break
		}
		fmt.Println(v)
	}
	<-flagb

}

func inse() {
	for i := 2; i <= 10000; i++ {
		num2Chan <- i
	}
	close(num2Chan)

}

func su2Test() {
	for {
		v, ok := <-num2Chan
		if !ok {
			break
		}
		flag := true
		for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
			if v%i == 0 {
				flag = false
				break

			}

		}
		if flag {
			res2Chan <- v
			//fmt.Println(v)
		}

	}
	flag2Chan <- true
}
