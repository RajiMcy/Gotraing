package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64
	go func() {

		for {

			atomic.AddUint64(&ops, 1)
			time.Sleep(time.Millisecond)
		}

	}()

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

}
