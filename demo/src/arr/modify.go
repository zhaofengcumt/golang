package main

import "fmt"

func modify(array [5]int) {
	array[0] = 0
	fmt.Println("modified value", array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("value", array)
}
