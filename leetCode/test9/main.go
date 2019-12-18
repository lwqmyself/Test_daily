package main

import "fmt"

func main() {
	num := 343
	fmt.Println(isPalindrome(num))
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	tmp := x
	for tmp != 0 {
		y = y*10 + tmp%10
		tmp = tmp / 10

	}
	fmt.Println(y)
	if x == y {
		return true
	} else {
		return false
	}

}
