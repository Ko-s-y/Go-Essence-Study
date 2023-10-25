package main

import (
	"fmt"
	"log"
)

type User struct {
	Name string
}

func main() {
	user, err := userName()
	if err == nil {
		fmt.Println(user.Name)
	} else {
		log.Println(err)
	}

	if user, err := userName(); err == nil {
		fmt.Println(user.Name)
	} else {
		log.Println(err)
	}
}

func userName() (*User, error) {
	return &User{Name: "John"}, nil
}
