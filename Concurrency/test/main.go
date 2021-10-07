package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {

		time.Sleep(1 * time.Second)

		c1 <- "Toy"
	}()
	select {
	case msg1 := <-c1:
		fmt.Println("Received:", msg1)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeedout")
	}

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "toy_image"
	}()
	select {
	case msg2 := <-c2:
		fmt.Println("Received:", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeedout")
	}

	go func() {
		time.Sleep(3 * time.Second)
		c3 <- "toy_video"
	}()

	select {

	case msg3 := <-c3:
		fmt.Println("Received:", msg3)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeedout")

	}

}
