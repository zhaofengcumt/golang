package main

import "fmt"

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = map[string]PersonInfo{
		"12345": PersonInfo{"12345", "Tom", "Room 203,..."},
		"22222": PersonInfo{"12345", "Tom", "Room 203,..."},
	}
	//personDB = make(map[string] PersonInfo)
	// 往这个map里插入几条数据
	//personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	//personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	// 从这个map查找键为"1234"的信息
	delete(personDB,"12345")
	person, ok := personDB["12345"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 12345.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}
