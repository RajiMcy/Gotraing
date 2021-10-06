package main

import (
	"fmt"
	"sync"
)

const turns = 1000000

func main() {

	num := 0
	done := make(chan bool)
	var mutex sync.Mutex

	go func() {

		for i := 0; i < turns; i++ {
			mutex.Lock()
			num++
			mutex.Unlock()
		}
		done <- true

	}()
	go func() {

		for i := 0; i < turns; i++ {
			mutex.Lock()
			num--
			mutex.Unlock()
		}
		done <- true

	}()

	<-done
	<-done

	fmt.Println(num)

}
