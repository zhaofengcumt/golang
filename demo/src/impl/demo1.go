package main

import(
	"fmt"
)

func main() {
	var v1 interface{} = new(Reader)
	switch v := v1.(type) {
	case Reader:
		fmt.Println(v)

	}
}
