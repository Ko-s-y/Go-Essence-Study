package main

import (
	"fmt"
)

func switch_case() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Print("two or ")
		fallthrough // two or three or four が返る
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}
}

func main() {
	switch_case()
}
