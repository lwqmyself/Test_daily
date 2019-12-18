package main

import "fmt"

func main() {
	num := 1324324
	fmt.Println(reverse(num))
	m := (1 << 31) - 1
	fmt.Println(m)
}
func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y

}
