package main

import (
	"fmt"
	"sync"
)

var result int

func main() {

	getnumber()

	fmt.Println(result)

}

func getnumber() {

	var wg sync.WaitGroup

	wg.Add(1)

	func(i int) {

		result = i
		wg.Done()
	}(5)

	wg.Wait()
}
