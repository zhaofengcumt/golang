package main

import "fmt"

func main() {

	str := "Hello world" // 字符串赋值
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
	for i, value := range str {
		fmt.Println(i, value)
	}

	i, j := 3, 4
	fmt.Println(i, j)
JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
		fmt.Println(j)
	}

	sum := 0
	for {
		sum++
		if sum > 100 {
			//	break
		}
	}

}
