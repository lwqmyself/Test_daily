package main

import (
	"fmt"
)

type Stu struct {
	No   int
	Next *Stu
}

func AddStu(num int) *Stu {
	helper := &Stu{}
	first := &Stu{}

	if num < 1 {
		fmt.Println("num太小")
		return first
	}
	for i := 1; i <= num; i++ {
		stu := &Stu{
			No:   i,
			Next: nil,
		}
		//需要一个辅助指针

		if i == 1 {
			first = stu
			helper = stu
			helper.Next = first
		} else {
			helper.Next = stu
			helper = stu
			helper.Next = first
		}

	}
	return first
}

//显示单项的环形链表

func show(first *Stu) {
	if first.Next == nil {
		fmt.Println("kong")
	}
	tmp := first
	for {
		fmt.Printf("编号 %d ", tmp.No)
		if tmp.Next == first {
			break
		}
		tmp = tmp.Next
	}
}

func Play(first *Stu, starNO, countNum int) {
	if first.Next == nil {
		fmt.Println("空的")
		return
	}
	tail := first
	for {
		if tail.Next == first {
			break

		}
		tail = tail.Next
	}
	for i := 1; i <= starNO-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	for {
		//开始数countNum次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next

		}
		fmt.Println("kill", first.No)
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
	fmt.Println("最后是", first.No)

}

func main() {
	first := AddStu(500)
	show(first)
	fmt.Println()
	Play(first, 20, 31)
}
