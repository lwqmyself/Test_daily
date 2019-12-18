package main

import "fmt"

type TTree struct {
	Name  string
	left  *TTree
	right *TTree
}

func main() {
	t1 := &TTree{
		Name:  "1",
		left:  nil,
		right: nil,
	}
	t2 := &TTree{
		Name:  "2",
		left:  nil,
		right: nil,
	}
	t3 := &TTree{
		Name:  "3",
		left:  nil,
		right: nil,
	}
	t4 := &TTree{
		Name:  "4",
		left:  nil,
		right: nil,
	}
	t5 := &TTree{
		Name:  "5",
		left:  nil,
		right: nil,
	}
	t6 := &TTree{
		Name:  "6",
		left:  nil,
		right: nil,
	}
	t7 := &TTree{
		Name:  "7",
		left:  nil,
		right: nil,
	}

	t1.left = t2
	t1.right = t3
	t2.left = t4
	t2.right = t5
	t3.left = t6
	t3.right = t7
	PreDFS(t1)

	BFS(t1)

}

func BFS(ht *TTree) {
	if ht == nil {
		return
	}
	var queue []*TTree
	queue = append(queue, ht)
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node.Name)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}

}

func PreDFS(t *TTree) {
	if t != nil {
		fmt.Println(t.Name)
		PreDFS(t.left)
		PreDFS(t.right)
	}
}
