package main

import "fmt"

func modify1(array []int) {
	array[0] = 0
	fmt.Println("modified value", array)
}

func main() {
	//var myarry = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var myslice = myarry[5:]
	var myslice = make([]int, 5, 10)
	var myslice1 = []int{1, 2, 0}
	fmt.Println(cap(myslice), len(myslice))
	myslice = append(myslice, 1, 2)
	myslice = append(myslice, myslice1...)
	fmt.Println(myslice)
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice1,slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1, slice2)
	var myslice2 = []int{1, 2, 0}
	fmt.Println(myslice2)
	modify1(myslice2)
	fmt.Println(myslice2)
}
