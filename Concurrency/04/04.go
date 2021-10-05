package main

import "fmt"

func main() {

	//unbufferd_channel

	numCh := make(chan int)
	go counter(numCh)
	//a := <-numCh
	//fmt.Println(a)
	//a = <-numCh
	//fmt.Println(a) we are getting all values with below for loop

	for a := range numCh {

		fmt.Println(a)
	}

}

func counter(out chan int) {

	for i := 0; i <= 3; i++ {

		out <- i

	}
	close(out) // without this 1st for loop will be waiting for the 2nd one to send more values
	// use above close statement to close the channel
}
