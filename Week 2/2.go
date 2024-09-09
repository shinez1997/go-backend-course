package main

import "fmt"

//Viết hàm có input là 1 string, trả về map có key là ký tự, và value là số lần xuất hiện của ký tự đó ở trong string

func get(s string) map[rune]int {
	res := make(map[rune]int)
	for _, x := range s {
		res[x]++
	}
	return res
}
func main() {
	s := "Golang Developer"
	fmt.Println(get(s))
}
