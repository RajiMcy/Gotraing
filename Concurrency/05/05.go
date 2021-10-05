package main

import "fmt"

func main() {

	// Generator -------> Printer

	numCh := make(chan int)
	done := make(chan struct{})
	go counter(numCh)
	go printer(numCh, done)

	<-done

}

func counter(out chan int) {

	for i := 0; i <= 3; i++ {

		out <- i

	}
	close(out)

}
func printer(in chan int, done chan struct{}) {
	for a := range in {

		fmt.Println(a)
	}
	done <- struct{}{}

}
