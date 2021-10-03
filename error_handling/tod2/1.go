package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {

		fmt.Println(err)

		return

	}

	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(string(b))
}
