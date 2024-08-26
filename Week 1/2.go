package main

import "fmt"

// 2. Viết chương trình nhập 1 string, in ra true nếu độ dài chuỗi chia hết cho 2, và false nếu ngược lại
func check(s string) bool {
	return len(s)%2 == 0
}
func main() {
	fmt.Println(check("abcd"))
	fmt.Println(check("abcde"))
	fmt.Println(check("abcdef"))

}
