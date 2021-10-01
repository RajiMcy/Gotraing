package main

import (
	"fmt"
)

func add(a1, a2 int) {

	res := a1 + a2

	fmt.Println("Result:", res)

}

func main() {
	defer fmt.Println("End")

	add(5, 5)
}
