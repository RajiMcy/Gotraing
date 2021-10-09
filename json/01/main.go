package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	res1D := response1{

		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

}
