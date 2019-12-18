package main

import "strings"

func main() {

}
func strStr(haystack string, needle string) int {
	if haystack == "" {
		return 0
	}
	i := strings.Index(haystack, needle)
	return i
}
