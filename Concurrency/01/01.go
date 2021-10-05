package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Begins")
	delayPrint()
	fmt.Println("Ends")
	time.Sleep(3 * time.Second)

}

func delayPrint() {

	time.Sleep(2 * time.Second)

	fmt.Println("Lazy Hi")

}
