package main

import (
	"fmt"
	"unicode"
)

const (
	a1 = iota
	a2
	a3
	a4 = 1
	a5 = iota
)

func main() {
	//char :="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//nums := "0123456789"

	str := "hello world,世界"
	str1 := []rune(str)
	fmt.Print(len(str), len(str1))
	count := 0
	for _, v := range str1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println()
	fmt.Print(a1, a2, a3, a4, a5)
	fmt.Println()
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Printf("%T", str1)

	a := make([]int, 0)
	a = append(a, 0)
	fmt.Println(a)
}
