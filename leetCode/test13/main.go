package main

import "fmt"

func main() {
	str := "D"
	fmt.Println(romanToInt(str))
}
func romanToInt(s string) int {
	res := 0
	m := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	count := 0
	i := 0
	for {
		if count > len(s)-2 {
			break
		}
		key := s[i : i+2]
		fmt.Println(key)
		if v, ok := m[key]; ok {
			res += v
			count += 2
			i += 2
		} else {
			key = s[i : i+1]
			res += m[key]
			count++
			i++
		}

	}
	fmt.Println(count)
	if len(s)-count == 1 {
		key := s[len(s)-1:]
		res += m[key]
	}
	return res

}
