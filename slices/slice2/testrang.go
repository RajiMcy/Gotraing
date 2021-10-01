package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 5}

	sum := 0

	for i, v := range nums {

		sum += v

		fmt.Println("i:", i)

	}

	fmt.Println("sum:", sum)

	for i, v := range nums {

		if v == 3 {

			fmt.Println("index:", i)

		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {

		fmt.Printf("%s ->%s\n", k, v)
	}

	for i, c := range "go" {

		fmt.Println(i, c)

	}

}
