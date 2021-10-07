package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := caller1()

	if err != nil {

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)

}

func caller1() (int, error) {

	err := caller2()
	if err != nil {
		return 0, fmt.Errorf("[caller1] error in calling caller2 %v", err)
	}

	return 1, nil

}

func caller2() error {

	err := caller3()
	if err != nil {
		return fmt.Errorf("[caller2] error in calling caller3 %v", err)
	}

	return nil

}

func caller3() error {

	return errors.New("failed")

}
