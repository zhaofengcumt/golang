package main

import (
	"intf"
	"fmt"
)

type Reader struct {
}

func (reader Reader) Read() int {
	fmt.Println("read")
	return 0
}

func main() {
	var reader intf.IReader = new(Reader)
	if _, ok := reader.(Reader); ok {
		fmt.Println("yes")
	}
	reader.Read()
	switch  value := reader.(type) {
	case intf.IReader:
		fmt.Println("yes",value)
	}

}
