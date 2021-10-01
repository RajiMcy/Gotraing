package main

import (
	"fmt"
)

type kitchen struct {
	numofplates int
}

type house struct {
	kitchen
	numofrooms int
}

func main() {

	h := house{kitchen{10}, 3}
	fmt.Println("rooms", h.numofrooms)
	fmt.Println("plates", h.numofplates)
	fmt.Println("kitchen content", h.kitchen)
}
