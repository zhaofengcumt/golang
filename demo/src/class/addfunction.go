package main

import "fmt"

type Integer int

func (pre *Integer) add(after Integer) {
	*pre += after
}

func main() {
	var value Integer = 3
	var value1 Integer = 4
	value.add(value1)
	fmt.Println(value)


}
