package main

import "fmt"

func main() {

	// Generator -------> squarer------> Printer

	numCh := make(chan int)
	done := make(chan bool)
	squareToPrint := make(chan int)
	go counter(numCh)
	go squarer(numCh, squareToPrint)
	// go squarer(input, output)
	go printer(squareToPrint, done)
	//go printer(input, output)

	<-done

}

func counter(out chan int) {

	for i := 0; i <= 3; i++ {

		out <- i

	}
	close(out)

}

func squarer(in chan int, out chan int) {

	for a := range in {

		out <- a * a

	}
	close(out)

}
func printer(in chan int, done chan bool) {
	for a := range in {

		fmt.Println(a)
	}
	done <- true

}
