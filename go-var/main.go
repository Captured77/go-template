package main

import (
	"fmt"
	"strconv"
)

/*
四种变量声明方式
*/

type user struct {
	name string
	age  string
	email
}

type email struct {
	uri    string
	name   string
	domain string
}

// 常量 （只读属性）
const length int = 10
const (
	A = iota
	B
	C
)

func main() {
	//fmt.Println(length)
	//fmt.Println(A, B, C)
	// test1()
	// test2()
	// fmt.Println(getExcelColumnName(1))
	// arr := [...]int{1, 2, 3}
	// test3(arr)
	// fmt.Println(arr)

	// var m map[string]int = make(map[string]int)
	// m["tom"] = 38

	// test4(m)

	// qcrao :=Person{age: 38}
	// fmt.Println(qcrao.howOld())
	// qcrao.growUp()
	// fmt.Println(qcrao.howOld())

	// stefno := &Person{age: 100}
	// fmt.Println(stefno.howOld())
	// stefno.growUp()
	// fmt.Println(stefno.howOld())

	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)

}

func test1() {
	// var a string; //方法一： 声明一个变量默认0
	// fmt.Printf("type of a =%T\n", a)
	// fmt.Println(reflect.TypeOf(a))

	// var t1 user;
	// fmt.Printf("type of t1 = %T\n", t1)
	// fmt.Println(reflect.TypeOf(t1))

	// var t2 user;
	// fmt.Printf("type of t2 = %T\n", t2)
	// fmt.Println(reflect.TypeOf(t2))

	fmt.Println("Hello world")
	f1 := "3.14"
	f2, _ := strconv.ParseFloat(f1, 32)
	fmt.Println(f2)

}

// go 语言中集合数据类型数据结构，数组、切片、map、list
func test2() {
	// 数组
	var a1 [3]string // 只有三个元素的数组类型
	var a2 [4]string // 只有四个元素的数组类型

	a1[0] = "go"
	a1[1] = "grpc"
	a1[2] = "gin"

	fmt.Printf("%v\n", a1)
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)

	// []string切片 和 [3]string 是不同的类型
	var s1 []string
	for i, v := range s1 {
		fmt.Println(i, v)
	}

	var a3 = [...]string{"go", "grpc", "gin"} // 放入多少个元素，就是多长
	var a4 [3]string
	a4[0] = "go"
	a4[1] = "grpc"
	a4[2] = "gin"
	for i, v := range a3 {
		fmt.Println(i, v)
	}

	for i, v := range a4 {
		fmt.Println(i, v)
	}

	if a3 == a4 { // 数组比较, 每个元素都会比较
		fmt.Println("a3 == a4")
	}

	// 多维数组
	var arr1 [3][4]string
	arr1[0] = [4]string{"go", "1h", "bobby"}
	arr1[1] = [4]string{"grpc", "2h", "alice"}
	arr1[2] = [4]string{"gin", "3h", "tom"}

	for i, v := range arr1 {
		for j := range v {
			fmt.Print(arr1[i][j] + " ")
		}
		fmt.Println()
	}

	for _, row := range arr1 {
		fmt.Println(row)
	}

	// NOTE: 切片
	var nums []int
	fmt.Printf("%v\n", nums)
	fmt.Println(len(nums), cap(nums))

	nums1 := append(nums, 1)
	fmt.Println(len(nums1), cap(nums1))

	nums2 := make([]int, 3, 5) // param 0: 类型  param 1: 长度，param 2: 容量
	fmt.Println(len(nums2), cap(nums2))

	fmt.Println("------------------------------")
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(slice), cap(slice))
	slice1 := slice[2:5]
	fmt.Println(len(slice1), cap(slice1))

	slice2 := slice1[2:6:7]
	fmt.Println(len(slice2), cap(slice2))

	slice2 = append(slice2, 100)
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice2 = append(slice2, 200)
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2, len(slice2), cap(slice2))

	fmt.Println("------------------------------------")
	slice1[2] = 20
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2, len(slice2), cap(slice2))

}

func getExcelColumnName(n int) string {
	column := ""
	for n > 0 {
		n-- // 调整索引，使其从0开始
		column = string(rune('A'+(n%26))) + column
		n /= 26
	}
	return column
}

func test3(s [3]int) {
	for i := range s {
		s[i] *= 2
	}
}

func test4(m map[string]int) {
	age1 := m["bobby"]
	fmt.Println(age1)

	age2, ok := m["bobby"]
	fmt.Println(age2, ok)
}

type Person struct {
	age int
}

func(p Person) howOld() int {
	return p.age
}

func(p *Person) growUp() {
	p.age++
}


type iGreeting interface{
	sayHello()
}
func sayHello(i iGreeting) {
	i.sayHello()
}

type Go struct{}
func(g Go) sayHello() {
	fmt.Println("I am Go")
}

type PHP struct{}
func(p PHP) sayHello() {
	fmt.Println("I am PHP")
}