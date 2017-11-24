package main

import "fmt"

func main() {
	str := "Hello world" // 字符串赋值
	ch := str[0]         // 取字符串的第一个字符
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
	str[0] = '1'
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
}
