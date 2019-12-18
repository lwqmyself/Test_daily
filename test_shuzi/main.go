package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := make([]int, 0)

	//var num string
	inputReader := bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')
	str1 := string(str)
	if err != nil {

	}
	str2 := strings.Trim(str1, "\r\n")
	fmt.Println(str2)
	//fmt.Scanf("%s", &num)
	nu := strings.Split(str2, " ")

	fmt.Println(nu)
	for i := 0; i < len(nu); i++ {
		tmp, _ := strconv.Atoi(nu[i])
		nums = append(nums, tmp)
		fmt.Println(tmp)
	}

	sort.Ints(nums)
	fmt.Println(nums)
}
