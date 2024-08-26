package main

import "fmt"

// 1. Viết ham nhập 2 cạnh của hình chữ nhật, in ra diện tích, chu vi
type Rectangle struct {
	H int
	W int
}

func (r *Rectangle) perimeter() int {
	return 2 * (r.H + r.W)
}
func (r *Rectangle) area() int {
	return r.H * r.W
}
func main() {
	var r Rectangle
	fmt.Scanf("%d%d", &r.H, &r.W)
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())
}
