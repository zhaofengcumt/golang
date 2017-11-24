package main

import "fmt"

func main() {
	/*	func(a, b int, z float64) bool {
			return a*b < int(z)
		}*/
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	func(x, y int) int {
		fmt.Println(x + y)
		return 2
	}(1, 2)

}
