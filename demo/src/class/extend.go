package main

import "fmt"

type Base struct {
	i int
}

func (base *Base) foo() { fmt.Println("baseFoo") }
func (base *Base) Bar() { fmt.Println("baseBar") }

type extend struct {
	i string
	Base
}

func main() {
	var myextend = new(extend)
	myextend.Bar()
	fmt.Println(myextend.i)
}
