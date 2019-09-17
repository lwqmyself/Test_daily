package main

import "fmt"

type HeroNode struct {
	num      int
	Name     string
	nickname string
	Next     *HeroNode
	Pre      *HeroNode
}

func InsertHeroNode(head, newHeroNode *HeroNode) {
	tmp := head
	flag := true
	for {
		if tmp.Next == nil {
			break
		} else if tmp.Next.num > newHeroNode.num {
			break
		} else if tmp.Next.num == newHeroNode.num {
			flag = false
			break
		}
		tmp = tmp.Next
	}
	if !flag {
		fmt.Println("节点已存在 num = ", newHeroNode.num)
		return
	} else {
		newHeroNode.Next = tmp.Next
		newHeroNode.Pre = tmp
		if tmp.Next != nil {
			tmp.Next.Pre = newHeroNode
		}
		tmp.Next = newHeroNode
	}
	//tmp.Next = newHeroNode
}
func ListHeroNOde(head *HeroNode) {
	tmp := head
	if tmp.Next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d] %s [%s]-->", tmp.Next.num, tmp.Next.Name, tmp.Next.nickname)
		tmp = tmp.Next
		if tmp.Next == nil {
			break
		}
	}
	fmt.Println()
}
func ListHeroNOde2(head *HeroNode) {
	//逆序打印
	tmp := head
	if tmp.Next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	for {
		fmt.Printf("[%d] %s [%s]-->", tmp.num, tmp.Name, tmp.nickname)
		tmp = tmp.Pre
		if tmp.Pre == nil {
			break
		}
	}
	fmt.Println()
}

func DelNode(head *HeroNode, id int) {
	tmp := head
	flag := false
	for {
		if tmp.Next == nil {
			break
		} else if tmp.Next.num == id {
			flag = true
			break
		}
		tmp = tmp.Next
	}
	if flag {
		tmp.Next = tmp.Next.Next
		if tmp.Next != nil {
			tmp.Next.Pre = tmp
		}
	} else {
		fmt.Println("id 不存在")
	}
}

func main() {
	//创建头节点
	head := &HeroNode{}

	hero1 := &HeroNode{
		num:      1,
		Name:     "宋江",
		nickname: "及时雨",
		Next:     nil,
	}
	hero2 := &HeroNode{
		num:      2,
		Name:     "卢俊义",
		nickname: "玉麒麟",
		Next:     nil,
	}
	hero3 := &HeroNode{
		num:      3,
		Name:     "林冲",
		nickname: "豹子头",
		Next:     nil,
	}

	//head.Next = hero1
	//hero1.Pre = head
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero2)

	ListHeroNOde(head)
	//DelNode(head, 2)
	//ListHeroNOde(head)
	DelNode(head, 2)
	ListHeroNOde2(head)
}
