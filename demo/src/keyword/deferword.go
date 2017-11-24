package main

import "fmt"

func test_defer(i int) int {
	return 0
}

func main() {
	defer test_defer(1)
	fmt.Println(test_defer(1))
}
