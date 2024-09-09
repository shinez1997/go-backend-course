package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	name, job string
}

func CreatePerson(s string) Person {
	tmp := strings.Split(s, ",")
	tmp[0] = strings.ToUpper(tmp[0])
	tmp[1] = strings.ToLower(tmp[1])
	return Person{tmp[0], tmp[1]}
}
func main() {
	f, e := os.Open("a.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var listPerson []Person
	for scanner.Scan() {
		listPerson = append(listPerson, CreatePerson(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(listPerson)
}
