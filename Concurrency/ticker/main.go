package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {

		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	go func() {
		time.Sleep(1600 * time.Millisecond)
		ticker.Stop()
		fmt.Println("ticker stopped")
	}()

	time.Sleep(2600 * time.Millisecond)
}
