package main

import "fmt"

type a12 int

type Student struct {
	Name string
	Age  int
}

func (aa a12) add() int {
	aa = aa + 11
	bb := int(aa)
	return bb
}

func (aa *a12) add2() int {
	*aa = *aa + 11
	bb := int(*aa)
	return bb
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name = [%v] Age = [%v]", stu.Name, stu.Age)
	return str
}

func main() {
	var res a12 = 10
	rea1 := res.add()
	fmt.Println(rea1)
	res.add2()
	fmt.Println(res)
	stu := Student{
		Name: "lwq",
		Age:  21,
	}
	fmt.Println(&stu)
}
