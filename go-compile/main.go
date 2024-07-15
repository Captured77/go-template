package main

import "fmt"

type Person interface {
	growUp()
}

type Student struct {
	age int
}

func (p Student) growUp() {
	p.age += 1
	return
}

// go build -gcflags="-S" main.go

func main() {
	var qcrao = Person(Student{age: 18})

	fmt.Println(qcrao)
}
