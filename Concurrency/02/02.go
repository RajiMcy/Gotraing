package main

import (
	"fmt"
	"time"
)

func main() {

	go delayPrint("first")
	go delayPrint("second")
	go delayPrint("third")

	time.Sleep(1 * time.Second)

}

func delayPrint(str string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s %d\n", str, i)

		time.Sleep(100 * time.Millisecond)

	}

}
