package main

import "fmt"

func recoverFullName() {

	if r := recover(); r != nil {

		fmt.Println("recovered from", r)

	}

}

func fullName(firstName *string, lastName *string) {

	defer recoverFullName()

	if firstName == nil {

		panic("runtime error:first name cannot be nil")

	}

	if lastName == nil {

		panic("runtime error:last name cannot be nil")

	}
	fmt.Printf("%s%s\n", *firstName, *lastName)
	fmt.Println("return normally from fullName")
}
func main() {

	defer fmt.Println("deffered call in main")
	firstName := "raji"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

}
