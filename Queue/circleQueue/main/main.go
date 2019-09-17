package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

//添加数据到队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//取
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列为空")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, err
}

//判满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//判元素个数
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

//显示队列
func (this *CircleQueue) showQueue() {
	size := this.Size()
	if this.IsEmpty() {
		fmt.Println("队列为空")
	}
	//
	tmpHead := this.head

	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d] = %d\n", tmpHead, this.array[tmpHead])
		tmpHead = (tmpHead + 1) % this.maxSize
	}

}

func main() {
	//创建队列
	cqueue := &CircleQueue{
		maxSize: 5,
		array:   [5]int{},
		head:    0,
		tail:    0,
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
			err := cqueue.Push(val)
			if err != nil {
				fmt.Println(err.Error())

			} else {
				fmt.Println("添加成功")
			}
		case "get":
			val, err := cqueue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(val)
		case "show":
			cqueue.showQueue()
		case "exit":
			os.Exit(0)

		}
	}
}
