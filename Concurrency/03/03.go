package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go delayPrint(done)
	<-done
	go delayPrint(done)
	<-done

}

func delayPrint(done chan bool) {

	time.Sleep(2 * time.Second)

	fmt.Println("Lazy Hi")
	done <- true

}
