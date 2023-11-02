package main

import (
	"fmt"
	"reflect"
)

func main() {
	type V int

	var v V = 123

	PrintDetail(v) // 定義したunderling typeを渡す
}

func PrintDetail(v any) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int/int32/int62型: ", v)
	case reflect.String:
		fmt.Println("strng型: ", v)
	default:
		fmt.Println("知らない型")
	}
}
