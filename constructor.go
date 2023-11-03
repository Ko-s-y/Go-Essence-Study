package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := NewUser("aaa", 12)
	fmt.Println(user.Name, user.Age)
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
