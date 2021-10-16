package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	//s := `Raji@242726`

	var s string

	fmt.Scanln(&s)

	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {

		fmt.Println(err)

	}

	fmt.Println(string(b))
	fmt.Println(s)
	//loginPwd := `Raji@24272`

	var loginPwd string
	fmt.Scanln(&loginPwd)

	err = bcrypt.CompareHashAndPassword(b, []byte(loginPwd))

	if err != nil {

		fmt.Println("You cannot login", err)
		return
	}

	fmt.Println("You can login")
}
