package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no = %d name = %s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no = %d name = %s\n", node.No, node.Name)

		InfixOrder(node.Right)
	}
}
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no = %d name = %s\n", node.No, node.Name)

	}
}

func main() {
	root := &Hero{
		No:    1,
		Name:  "a",
		Left:  nil,
		Right: nil,
	}
	left1 := &Hero{
		No:    2,
		Name:  "b",
		Left:  nil,
		Right: nil,
	}
	right1 := &Hero{
		No:    3,
		Name:  "c",
		Left:  nil,
		Right: nil,
	}
	root.Left = left1
	root.Right = right1
	right2 := &Hero{
		No:    5,
		Name:  "e",
		Left:  nil,
		Right: nil,
	}
	left2 := &Hero{
		No:    4,
		Name:  "d",
		Left:  nil,
		Right: nil,
	}
	left1.Left = left2
	right1.Left = right2
	//PreOrder(root)
	InfixOrder(root)
	PostOrder(root)
}
