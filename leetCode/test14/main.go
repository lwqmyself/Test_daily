package main

import (
	"fmt"
	"strings"
)

/*编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	//return ""
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}
	tmp := strs[0][0:1]
	count := 0
loop:
	for {
		count++
		if count > len(strs[0]) {
			break
		}
		tmp = strs[0][0:count]
		fmt.Println(tmp)
		for i := 0; i < len(strs); i++ {
			if strings.HasPrefix(strs[i], tmp) {
				continue
			} else {
				count--
				fmt.Println(count)
				break loop
			}
		}

	}

	//fmt.Println(tmp)
	return strs[0][0:count]
}
