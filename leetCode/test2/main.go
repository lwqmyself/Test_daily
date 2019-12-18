package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersWrong(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp1, tmp2 *ListNode

	var result []ListNode

	tmp1 = l1
	tmp2 = l2
	num1 := 0
	num2 := 0
	count := 0
	for {
		num1 += tmp1.Val * int(math.Pow10(count))
		count++
		if tmp1.Next == nil {
			break
		}
		tmp1 = tmp1.Next

		//fmt.Println(tmp1.Val)
	}
	count2 := 0
	for {
		num2 += tmp2.Val * int(math.Pow10(count2))
		count2++

		if tmp2.Next == nil {
			break
		}
		tmp2 = tmp2.Next

		//fmt.Println(num2,tmp2.Val,int(math.Pow10(count2)),count2)
	}

	res := num1 + num2
	res1 := strconv.Itoa(res)
	fmt.Println(num1, num2)
	for i := len(res1) - 1; i >= 0; i-- {
		tmp, _ := strconv.Atoi(res1[i : i+1])
		fmt.Println(tmp)
		result = append(result, ListNode{Val: tmp})
	}
	//fmt.Println(len(result))
	for i := 0; i < len(result)-1; i++ {
		result[i].Next = &result[i+1]
	}

	return &result[0]
}
func main() {
	head1 := &ListNode{}
	head2 := &ListNode{}
	n1 := []int{2, 4, 3}
	n2 := []int{5, 6, 4}
	for i := 0; i < 3; i++ {
		tmp := &ListNode{}
		tmp.Val = n1[i]
		InsertHeroNode(head1, tmp)
	}
	for i := 0; i < 3; i++ {
		tmp := &ListNode{}
		tmp.Val = n2[i]
		InsertHeroNode(head2, tmp)
	}

	ListHeroNOde(addTwoNumbersTest(head1, head2))

}
func InsertHeroNode(head, newHeroNode *ListNode) {
	tmp := head

	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}

	//newHeroNode.Next = tmp.Next
	tmp.Next = newHeroNode

	//tmp.Next = newHeroNode
}
func ListHeroNOde(head *ListNode) {
	tmp := head
	if tmp.Next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d]-->", tmp.Val)
		if tmp.Next == nil {
			break
		}

		tmp = tmp.Next

	}
	fmt.Println()
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义一个 i 表示 两数相加是否大于10, 定义一个s表示两数两加之和
	var i, s int
	// 定义一个头节点
	res := &ListNode{Val: 0}
	// 定义一个当前节点
	now := res

	// 无限循环,内部加条件跳出
	for true {
		// 如果 i >0 表示前面的数之和大于10了,所以当前的和要加1,否则不加
		if i > 0 {
			s = l1.Val + l2.Val + 1
		} else {
			s = l1.Val + l2.Val
		}
		// 如果 两数之和大于10,那么该位的值应该是 s - 10 ,否则就是 和 本身,并设置 i 标记
		if s >= 10 {
			now.Next = &ListNode{Val: s - 10}
			i = 1
		} else {
			now.Next = &ListNode{Val: s}
			i = 0
		}
		// 将 当前节点移动到下一指针
		now = now.Next

		// 当l1 和 l2 都移动到最后了,准备跳出循环
		if l1.Next == nil && l2.Next == nil {
			// 如果l1 和 l2 最后的和大于10,即 i == 1, 那么后面还需要添加一个 1 (小学数学问题,不用解释为什么是1吧)
			if i == 1 {
				now.Next = &ListNode{Val: 1}
			}
			// 跳出循环
			break
		}
		// 如果 l1走到了最后,可能l2 还没结束,所以把l1的当前节点值设置成0,继续跟l2相加,否则移动到下一指针
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		// 如果 l2走到了最后,可能l1 还没结束,所以把l2的当前节点值设置成0,继续跟l1相加,否则移动到下一指针
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}

	// 返回头节点的下一个节点指针,因为头节点是我设置成0的
	//ListHeroNOde(res)
	return res.Next

}

func addTwoNumbersTest(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	res := &ListNode{}
	flag := false
	tmp := res
	for true {
		if flag {
			sum = l1.Val + l2.Val + 1
		} else {
			sum = l1.Val + l2.Val
		}
		if sum >= 10 {
			tmp.Next = &ListNode{Val: sum - 10}
			flag = true
		} else {
			tmp.Next = &ListNode{Val: sum}
			flag = false
		}
		if l1.Next == nil && l2.Next == nil {
			if flag {
				tmp.Next = &ListNode{Val: 1}
			}
			break
		}
		tmp = tmp.Next
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}

	}
	return res.Next
}
