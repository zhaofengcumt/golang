package main

import (
	"fmt"
)

const (
	name = "kobe"
	age  = 25
	zero = iota
	one  = iota
)
const (
	ling = iota
	yi   = iota
)

const (
	Sunday       = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays  // 这个常量没有导出
)

func main() {
	fmt.Println(name, age)
	fmt.Println(zero, one)
	fmt.Println(ling, yi)
	fmt.Println(Sunday, Monday)
	fmt.Println(Saturday)
}
