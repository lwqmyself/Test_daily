package main

import "fmt"

type CateNode struct {
	no   int
	name string
	next *CateNode
}

func insertCat(head *CateNode, newCatNode *CateNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode.no, "已加入")
		return
	}
	tmp := head
	for {
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}
	tmp.next = newCatNode
	newCatNode.next = head
}
func List(head *CateNode) {
	fmt.Println("-----------------------")
	tmp := head
	if tmp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("猫的信息为 [id = %d name = %s]\n", tmp.no, tmp.name)
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}
}

func Del(head *CateNode, id int) *CateNode {
	tmp := head
	helper := head
	if tmp.next == nil {
		fmt.Println("kong")
		return head
	}
	//如果只有一个节点
	if tmp.next == head {
		tmp.next = nil
		return head
	}

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	//两个以上节点
	for {
		if tmp.next == head {
			break
		}
		if tmp.no == id {
			//找到
			if tmp == head {
				head = head.next
			}
			helper.next = tmp.next
			fmt.Printf("%d kill\n", id)
			flag = false
			break
		}
		tmp = tmp.next
		helper = helper.next
	}
	if flag {
		//说明在上面没找到
		if tmp.no == id {
			//找到
			helper.next = tmp.next
			fmt.Printf("%d kill\n", id)

		} else {
			fmt.Println("未找到")
		}
	}
	return head

}
func main() {
	head := &CateNode{
		no:   0,
		name: "",
		next: nil,
	}
	cat1 := &CateNode{
		no:   1,
		name: "white",
		next: nil,
	}
	cat2 := &CateNode{
		no:   2,
		name: "white",
		next: nil,
	}
	cat3 := &CateNode{
		no:   3,
		name: "white",
		next: nil,
	}
	cat4 := &CateNode{
		no:   4,
		name: "white",
		next: nil,
	}
	cat5 := &CateNode{
		no:   5,
		name: "white",
		next: nil,
	}

	insertCat(head, cat1)
	insertCat(head, cat2)
	insertCat(head, cat3)
	insertCat(head, cat4)
	insertCat(head, cat5)

	List(head)
	head = Del(head, 1)
	List(head)
}
