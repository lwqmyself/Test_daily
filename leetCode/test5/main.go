package main

import "fmt"

func main() {
	str := "aaaa"
	fmt.Println(longestPalindrome1111(str))
}
func longestPalindrome1111(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	maxi := 0
	maxj := 0
	maxCount := 0
	maxCountiii := 0
	for i := 0; i < l; i++ {
		//fmt.Print(s[i : i+1]," ")
		if i == 0 {
			continue
		}
		count := 0
		for j := 1; i-j >= 0 && i+j <= l-1; j++ {
			fmt.Println(s[i-j:i-j+1], s[i+j:i+j+1])
			if s[i-j:i-j+1] == s[i+j:i+j+1] {
				count++
			} else {
				break
			}
		}
		fmt.Println("count =", count)
		if count > maxj {
			maxj = count
			maxi = i
		}
		count2222 := 0
		for k := 1; i-k >= 0 && i+k <= l; k++ {
			if s[i-k] == s[i+k-1] {
				count2222++
			} else {
				break
			}
		}
		if count2222 > maxCount {
			maxCount = count2222
			maxCountiii = i
		}
	}
	//fmt.Println("maxCount =",maxCount)
	if maxj == 0 && maxCount == 0 {
		return s[0:1]
	}
	maxjck := 1 + maxj*2
	maxaxa := maxCount * 2
	fmt.Println("maxCount =", maxCount, "maxj =", maxj)
	if maxjck > maxaxa {
		return s[maxi-maxj : maxi+maxj+1]
	} else {
		return s[maxCountiii-maxCount : maxCountiii+maxCount]
	}

}
func longestPalindrome(s string) string {
	res1i, res1j, res2i, res2j := 0, 0, 0, 0
	for i, j := 0, 1; i < len(s) && j <= len(s); {
		if compute(i, j, s) {
			res1i, res1j = i, j
			if i == 0 {
				j += 2
			} else {
				i--
				j++
			}
		} else {
			i++
			j++
		}
	}
	for i, j := 0, 0; i < len(s) && j <= len(s); {
		if compute(i, j, s) {
			res2i, res2j = i, j
			if i == 0 {
				j += 2
			} else {
				i--
				j++
			}
		} else {
			i++
			j++
		}
	}
	if res1j-res1i > res2j-res2i {
		return s[res1i:res1j]
	} else {
		return s[res2i:res2j]
	}

}

func compute(i int, j int, s string) bool {
	if j-i > 1 {
		if s[i] != s[j-1] {
			return false
		}
		return compute(i+1, j-1, s)
	} else {
		return true
	}
}
