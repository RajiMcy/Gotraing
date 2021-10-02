package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var inside int

var dartCount int

var mtx sync.Mutex

var wg sync.WaitGroup

func pi() float32 {

	//increase dart count here to get more accurate results.
	//fmt.Println("Enter the number of darts :")
	//fmt.Scan(&dartCount)
	dartCount = 1000000

	//adding total number of go routines to waitgroups
	wg.Add(dartCount + 1)

	for i := 0; i <= dartCount; i++ {

		go func() {
			mtx.Lock()
			x := rand.Float32()
			y := rand.Float32()

			if x*x+y*y < 1 {
				inside++
			}
			mtx.Unlock()
			wg.Done()

		}()
	}

	wg.Wait()

	ratio := float32(inside) / float32(dartCount)
	return ratio * 4

}

func main() {

	fmt.Printf("value of PI after dartcount:%d is : %.2f\n", dartCount, pi())

}
