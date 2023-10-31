package main

import "fmt"

type Member struct {
	Name string
	Age  int
}

func main() {
	var member Member
	member.Name = "Bob"
	member.Age = 18

	fmt.Println(member)

	mem := Member{
		Name: "Michel",
		Age:  18,
	}

	fmt.Println(mem)
}
