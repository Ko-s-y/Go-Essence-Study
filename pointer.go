package main

import "fmt"

func main() {
	v := 1
	p := &v
	*p = 2
	fmt.Println(v)
}
