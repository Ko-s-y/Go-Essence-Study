package main

import "fmt"

const (
	Apple = iota
	Orange
	Banana
)

func main() {
	fmt.Println(Apple)  // 0
	fmt.Println(Orange) // 1
	fmt.Println(Banana) // 2
}
