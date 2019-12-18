package main

import "fmt"

type A struct {
	Id   int
	name string
}
type B struct {
	Id   int
	name string
}

func main() {
	//fmt.Println(A{} ==B{})
	var a1 A
	var a2 A
	b := B(a2)
	fmt.Println(a1 == a2)
}
