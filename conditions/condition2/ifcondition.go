package main

import (
	"fmt"
)

func main() {

	if num := 11; num < 0 {
		fmt.Println(num, "is negative")

	} else if num < 10 {

		fmt.Println(num, "has 1 digit")

	} else {

		fmt.Println(num, "has multiple digit")

	}

}
