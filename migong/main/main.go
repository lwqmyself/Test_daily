package main

import "fmt"

func SetWay(myMap *[7][7]int, i, j int) bool {
	if myMap[5][4] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 {
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) {
				fmt.Println("下")
				return true
			} else if SetWay(myMap, i, j+1) {
				fmt.Println("右")
				return true
			} else if SetWay(myMap, i-1, j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else {
				fmt.Println("wwww")
				myMap[i][j] = 3
				return false
			}

		} else {
			return false
		}
	}
}

func main() {
	var map1 [7][7]int
	for i := 0; i < 7; i++ {
		map1[0][i] = 1
		map1[6][i] = 1
	}
	for i := 0; i < 7; i++ {
		map1[i][0] = 1
		map1[i][6] = 1
	}
	//map1[5][4] = 2
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(map1[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	map1[3][1] = 1
	map1[3][2] = 1
	map1[3][3] = 1
	map1[3][4] = 1
	SetWay(&map1, 1, 1)
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(map1[i][j], " ")
		}
		fmt.Println()
	}
}
