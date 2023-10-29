package main

import (
	"fmt"
	"log"
)

type User struct {
	Name string
}

func userName() (*User, error) {
	return &User{Name: "John"}, nil
}

// go run if_else.go
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
