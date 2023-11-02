package main

import "fmt"

type V int

var v V = 123

func main() {
	any_test()
	PrintDetail(42)        // int型を渡す
	PrintDetail(int32(42)) // int32型を渡す
	PrintDetail(int64(42)) // int64型を渡す
	PrintDetail("Hello")   // string型を渡す
	PrintDetail(v)         // 定義したunderling typeを渡す
}

func any_test() {
	var v any // any は interface{}

	v = "hello world"
	fmt.Println(v)
	v = 1
	fmt.Println(v)

	s, ok := v.(string)
	if !ok {
		fmt.Println("vはstringではない")
	} else {
		fmt.Printf("vは文字列 %q です", s)
	}

	v = 1
	fmt.Println(v)
	v = "hello world"
	fmt.Println(v)

	if s, ok := v.(string); !ok {
		fmt.Println("vはstringではない")
	} else {
		fmt.Printf("vは文字列 %q です", s)
	}
}

func PrintDetail(v interface{}) {
	switch t := v.(type) {
	case int, int32, int64:
		fmt.Println("int/int32/int64 型", t)
	case string:
		fmt.Println("string 型:", t)
	default:
		fmt.Println("知らない世界")
	}
}
