package main

import "fmt"

type Cat struct {
	name  string
	color string
	age   int
}

func main() {
	map_a := map[int]string{1: "a"}
	var map_b map[string]map[string]string
	map_b = make(map[string]map[string]string)
	//map1 := make(map_a)
	//map_a[1] = "map_a 1"
	map_b["b"] = map[string]string{"c": "d"}
	map_b["b"]["c"] = "e"
	fmt.Println(map_a)
	fmt.Println(map_b)
	var cat1 Cat
	cat1.age = 10
	cat1.color = "white"
	fmt.Printf("%p\n", &cat1)
	fmt.Println(&cat1.color)
	cat1.Rec()
}

func (cat Cat) Rec() {
	fmt.Println("22")
}
