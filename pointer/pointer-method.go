package main

import "fmt"

// User struct
type User struct {
	Name string
}

func (user *User) addLastName(lastName string) {
	user.Name = user.Name + " " + lastName
}

func pointerMethod() {
	user := User{"Odia"}
	fmt.Println(user.Name)

	user.addLastName("Pratama")
	fmt.Println(user.Name)
}
