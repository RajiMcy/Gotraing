package main

import (
	"fmt"
)

var result int

func fibo(index int) int {

	if index <= 1 {
		return index

	}

	return fibo(index-1) + fibo(index-2)

}

func main() {

	n := 5

	if n == 0 {

		fmt.Println(-1, "(error)")

	} else if n == 40 {

		fmt.Println(-1, "(error)")

	} else {

		for i := 1; i <= n; i++ {

			result = (fibo(i))

		}

	}
	fmt.Println(result)
}
