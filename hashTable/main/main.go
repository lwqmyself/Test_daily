package main

import (
	"fmt"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head *Emp
}

//添加时，保证从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head //辅助指针
	var pre *Emp = nil
	//如果当前emplink为空
	if cur == nil {
		this.Head = emp
		return
	}
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next

		} else {
			break
		}
	}

	pre.Next = emp
	emp.Next = cur

}

type HashTable struct {
	LinkArr [7]EmpLink
}

//
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)

}

func (this *HashTable) HashFun(id int) int {
	return id % 7 //对应链表的下标
}
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表 %d 为空\n", no)
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表 %d -- 雇员%d = %d 名字 = %s", no, cur.Id, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
		fmt.Println()
	}
}

func (this *HashTable) showAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func main() {

	emp := &Emp{
		Id:   5,
		Name: "qewrwerw",
	}
	var hashT HashTable
	hashT.Insert(emp)
	hashT.showAll()
}
