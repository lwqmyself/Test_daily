package main

import "fmt"

func main() {
	x := 321
	y := 100

	fmt.Println(rev(rev(x) + rev(y)))
}
func rev(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10

		x /= 10
	}
	return y

}
