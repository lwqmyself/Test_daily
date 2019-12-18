package main

import "fmt"

func main() {

	same := "]"

	fmt.Println(isValid(same))

}

func isValid(s string) bool {
	st := initStack()
	// 后括号映射表
	frontBracket := map[string]string{")": "(", "]": "[", "}": "{"}
	for _, val := range s {
		if val == '(' || val == '[' || val == '{' {
			st.Push(byte(val))
		} else if len(st.item) != 0 {
			b := string(st.Pop().(byte))
			fmt.Println(b)
			if b == frontBracket[string(val)] {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if len(st.item) == 0 {
		return true
	} else {
		return false
	}

}

type stack struct {
	item []interface{}
}

func initStack() *stack {

	st := stack{}
	st.New()

	return &st
}

// 创建栈
func (s *stack) New() *stack {
	s.item = make([]interface{}, 0)
	return s
}

func (s *stack) Pop() interface{} {
	if len(s.item) == 0 {
		return nil
	}
	item := s.item[len(s.item)-1]
	s.item = s.item[0 : len(s.item)-1]

	return item
}

func (s *stack) Push(t byte) {
	s.item = append(s.item, t)
}
