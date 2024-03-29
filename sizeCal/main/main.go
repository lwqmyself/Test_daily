package main

import (
	"Test_daily_2333/sizeCal/stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	print("请输入正确的数学表达式: ")
	var stat string
	reader := bufio.NewReader(os.Stdin)
	stat, _ = reader.ReadString('\n')
	stat = strings.TrimSpace(stat)
	postfix := infix2ToPostfix(stat)
	fmt.Printf("后缀表达式：%s\n", postfix)
	fmt.Printf("计算结果: %d\n", calculate(postfix))
}

func calculate(postfix []string) int {
	stack := stack.ItemStack{}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		_, err := strconv.Atoi(postfix[i])
		//fmt.Println(postfix[i])
		if err == nil {
			stack.Push(nextChar)
		} else {
			// 操作符：取出两个数字计算值，再将结果压栈
			num1, _ := strconv.Atoi(stack.Pop().(string))
			num2, _ := strconv.Atoi(stack.Pop().(string))
			fmt.Println(num2, num1, nextChar)
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num2 + num1))
			case "-":
				stack.Push(strconv.Itoa(num2 - num1))
			case "*":
				stack.Push(strconv.Itoa(num2 * num1))
			case "/":
				stack.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top().(string))
	return result
}

func infix2ToPostfix(exp string) []string {
	stack := stack.ItemStack{}
	var postfix []string
	expLen := len(exp)

	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])
		stack.Show()
		switch char {
		case " ":
			continue
		case "(":
			// 左括号直接入栈
			stack.Push("(")
		case ")":
			// 右括号则弹出元素直到遇到左括号
			for !stack.IsEmpty() {
				//取栈顶
				preChar := stack.Top().(string)
				if preChar == "(" {
					stack.Pop() // 弹出 "("
					break
				}
				//加入切片
				postfix = append(postfix, preChar)
				//弹出
				stack.Pop()
			}

			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}

			postfix = append(postfix, digit)
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1

		default:
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for !stack.IsEmpty() {
				top := stack.Top().(string)

				if top == "(" || isLower(top, char) {
					fmt.Println(isLower(top, char))
					break
				}
				postfix = append(postfix, top)
				fmt.Println(postfix)
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}

	// 栈不空则全部输出
	for !stack.IsEmpty() {
		postfix = append(postfix, stack.Pop().(string))
		fmt.Println(postfix)
	}

	return postfix
}

func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}
