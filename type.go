package main

import "fmt"

type MyString string

func main() {
	var m MyString
	m = "foo"
	fmt.Println(m)

	s := string(m) // MyString型からstring型に戻す
	fmt.Println(s)
}
