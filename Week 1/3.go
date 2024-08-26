package main

import (
	"fmt"
	"slices"
)

// 3. Viết chương trình nhập một slice số, in ra tổng, số lớn nhất, số nhỏ nhất, trung bình cộng, slice đã được sắp xếp

func main() {
	var a []int
	var n, sum int
	var avg float32
	fmt.Scanf("%d", &n)
	for range n {
		var x int
		fmt.Scanf("%d", &x)
		a = append(a, x)
		sum += x
	}
	maxe, mine := a[0], a[0]
	for _, x := range a {
		maxe = max(maxe, x)
		mine = min(mine, x)
	}
	avg = (float32)(sum / n)
	slices.Sort(a)
	fmt.Println("Sum : ", sum)
	fmt.Println("Avg : ", avg)
	fmt.Println("Min : ", mine)
	fmt.Println("Max : ", maxe)
	fmt.Println("Sorted: ", a)

}
