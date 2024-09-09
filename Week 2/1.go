package main

import (
	"fmt"
)

// Viết hàm tạo ra 1 struct về 1 người (gồm tên, nghề nghiệp, năm sinh) , và struct có method tính tuổi,  method kiểm tra người ấy có hợp với nghề của mình không ( năm sinh chia hết cho số chữ trong tên)

type Person struct {
	year      int
	name, job string
}

func (p *Person) age() int {
	return 2024 - p.year
}

func (p *Person) check() {
	if p.year%len(p.name) == 0 {
		fmt.Println("Nghe hop tuoi")
	} else {
		fmt.Println("Nghe khong hop tuoi")
	}
}
func main() {
	a := Person{
		year: 1997,
		name: "S",
		job:  "Developer",
	}
	a.check()
}
