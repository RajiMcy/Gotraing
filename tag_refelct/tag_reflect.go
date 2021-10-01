package main

import (
	"fmt"
	"reflect"
)

const tagName = "validate"

type person struct {
	id    int    `validate:"_"`
	name  string `validate:"presence,min=2,max=32"`
	email string `validate:"email,required"`
}

func main() {
	user := person{

		id:    1,
		name:  "Raji",
		email: "raji123@gamil.com",
	}

	t := reflect.TypeOf(user)
	fmt.Println(t)
	fmt.Println("Type", t.Name())
	fmt.Println("Kind", t.Kind())

	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		fmt.Printf("%d.%v (%v),tag:'%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}

}
