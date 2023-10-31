package main

import "fmt"

func main() {
	string_join()
	string_converter()
	rune_string_converter()
	change_array_to_slice()
	string_multipul_lines()
}

func string_join() {
	name := "John"
	greet := "Hello "
	greet += name
	fmt.Println(greet)           // => Hello John
	fmt.Printf("%c\n", greet[0]) // => H
	fmt.Println(greet[0])        // => 72 HのUnicodeポインタ
}

func string_converter() {
	greet := "Hello"
	b := []byte(greet) // バイト列に変換
	b[0] = 'h'         // バイトスライスの最初の要素を変更
	greet = string(b)  // イミュータブルなので再代入が必要
	fmt.Println(greet)
}

func rune_string_converter() {
	greet := "こんにちわ世界"
	rune_greet := []rune(greet)
	rune_greet[4] = 'は'        // わをはに置換
	greet = string(rune_greet) // イミュータブルなので再代入が必要
	fmt.Println(greet)
}

func change_array_to_slice() {
	var b [4]byte
	b2 := b[:] // b2はスライス
	fmt.Println(b2)
}

func string_multipul_lines() {
	content := `こんにちは
今日はいい天気ですね。
明日も晴れるといいですね。`
	fmt.Println(content)
}
