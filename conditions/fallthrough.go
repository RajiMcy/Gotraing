package main

import (
	"fmt"
)

func main() {
	i := 15
	switch {

	case i <= 10:
		fmt.Println("less than or equal to 10")
		fallthrough

	case i >= 20:
		fmt.Println("greater than or equal to 20")

	case i >= 30:
		fmt.Println("greater than or equal to 30")

	default:
		fmt.Println("greater than 20")

	}
}
