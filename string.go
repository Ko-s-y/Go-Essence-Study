package main

import "fmt"

// go run string.go
func main() {
	string_join()
	string_converter()
}

func string_join() {
	name := "John"
	greet := "Hello "
	greet += name
	fmt.Println(greet)         // => Hello John
	fmt.Println(greet[0])      // => 72 HのUnicodeポインタ
	fmt.Printf("%c", greet[0]) // => H
}

func string_converter() {
	greet := "Hello"
	b := []byte(greet) // バイト列に変換
	b[0] = 'h'         // Helloをhelloに変換
	greet = string(b)  // stringはイミュータブルなため再代入が必要
	fmt.Println(greet)
}
