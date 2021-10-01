package main

import (
	"fmt"
)

type purchase interface {
	sell()
}

type display interface {
	show()
}

type salesman interface {
	purchase
	display
}

type game struct {
	name           string
	price          int
	gamecollection []string
}

func (t game) sell() {

	fmt.Println("-----------------------------")
	fmt.Println("Name:", t.name)
	fmt.Println("Price:", t.price)
	fmt.Println("-----------------------------")

}
func (t game) show() {

	fmt.Println("The games available are:")
	for _, name := range t.gamecollection {
		fmt.Println(name)
	}
}

func shop(s salesman) {

	fmt.Println(s)
	s.sell()
	s.show()
}

func main() {

	collection := []string{"yz", "pubg", "candycrush", "ABC"}
	game1 := game{"ert", 500, collection}
	shop(game1)
}
