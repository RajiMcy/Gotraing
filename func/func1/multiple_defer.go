package main

import (
	"fmt"
)

func add(a1, a2 int) {

	res := a1 + a2

	fmt.Println("Result:", res)

}

func main() {

	fmt.Println("start")
	defer fmt.Println("End")

	//Executes in LIFO order

	defer add(5, 5)
	add(10, 10)
	defer add(15, 20)
}
