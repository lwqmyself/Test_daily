package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("队列已满")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//取
func (this *Queue) GetQueue() (val int, err error) {

	if this.rear == this.front {

		return -1, errors.New("空队列")

	}
	this.front++
	val = this.array[this.front]

	return val, err
}

func (this *Queue) showQueue() {
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, this.array[i])
	}
}

func main() {
	//创建队列
	queue := &Queue{
		maxSize: 5,
		array:   [5]int{},
		front:   0,
		rear:    0,
	}
	var (
		key string
		val int
	)
	for {
		fmt.Println("1 输入add添加")
		fmt.Println("2 输入get获取")
		fmt.Println("3 输入show显示队列")
		fmt.Println("4 输入exit离开")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队的数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())

			} else {
				fmt.Println("添加成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(val)
		case "show":
			queue.showQueue()
		case "exit":
			os.Exit(0)

		}
	}
}
