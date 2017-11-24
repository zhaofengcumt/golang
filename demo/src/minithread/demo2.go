package main

import (
	"time"
	"fmt"
)

func main() {

	/*	ch := make(chan int, 1)
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
			i := <-ch
			fmt.Println("Value received:", i)
		}*/

	timeout := make(chan bool, 1)
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(3e9) // 等待1秒钟
		timeout <- true
	}()
	// 然后我们把timeout这个channel利用起来
	select {
	case <-ch:
		// 从ch中读取到数据
	case <-timeout:
		fmt.Println("read ok")
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
	}

}
