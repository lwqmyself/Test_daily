package main

import (
	"fmt"
)

func main() {

	var n int
	str := ""
	xx := make([]byte, 0)
	count := make([]int, 0)
	res := make([]int, 0)
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s", &str)
	for i := 0; i < len(str); i++ {
		xx = append(xx, str[i])
	}
	//fmt.Println(len(xx))
	for i := 0; i < len(xx); i++ {
		if xx[i] == 'O' {
			count = append(count, -1)

		} else {
			count = append(count, 0)
		}
	}
	//nextCount := 0
	//preCount := 0
	for i := 0; i < len(count); i++ {
		nexttmp := 0
		pretmp := 0
		nextflag := false
		preflag := false
		if count[i] == 0 {
			for {

				if i+nexttmp > len(count)-1 {

					break
				} else if count[i+nexttmp] == -1 {
					nextflag = true
					break
				}
				nexttmp++
			}
			if nextflag == false {
				nexttmp = 0
			}
			for {

				if i-pretmp < 0 {
					break
				} else if count[i-pretmp] == -1 {
					preflag = true
					break
				}
				pretmp++
			}
			if preflag == false {
				pretmp = 0
			}

		}

		//fmt.Println("nextmp",nexttmp,"pre ",pretmp)
		if pretmp > nexttmp {
			if nexttmp == 0 {
				res = append(res, pretmp)
			} else {
				res = append(res, nexttmp)
			}

		} else {
			if pretmp == 0 {
				res = append(res, nexttmp)
			} else {
				res = append(res, pretmp)
			}

		}
	}
	for _, v := range res {
		fmt.Printf("%d ", v)
	}

}
