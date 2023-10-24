package main

import "fmt"

type Fruit int
type Animal int

const (
	Apple  Fruit = iota //Fruit(0)
	Orange              //Fruit(1)
	Banana              //Fruit(2)
)

const (
	Monky    Animal = iota //Animal(0)
	Elephant               //Animal(0)
	Pig                    //Animal(0)
)

func iota_study() {
	var fruit Fruit = Apple
	fmt.Println(fruit)

	// 下記はコンパイルエラーが発生する
	// fruit = Elephant
	// fmt.Println(fruit)
	// # command-line-arguments
	// ./main.go:24:10: cannot use Elephant (constant 1 of type Animal) as Fruit value in assignment
}
