package main

import (
	"fmt"
)

func do(i interface{}) {

	switch v := i.(type) {

	case int:

		fmt.Printf("Twice %v is %v\n", v, v*2)

	case string:

		fmt.Println(len(v))

	default:

		fmt.Printf("IDk the type%T!", v)

	}

}
func main() {
	do(21)
	do("hello")
	do(true)
}
