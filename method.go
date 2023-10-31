package main

import "fmt"

type Value int // カスタム型 Value を定義

func main() {
	v := Value(1)
	fmt.Println(v)
	v = v.Add(Value(2)) // インスタンス v をメソッド Add で更新
	fmt.Println(v)
}

func (v Value) Add(n Value) Value {
	return v + n
}
