package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer_prints()
	open_file()
	do_something1()
	do_something2()
}

func open_file() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))
}

func defer_prints() {
	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func do_something1() {
	var n int = 1
	defer fmt.Println(n)
	n = 2

	// 下記と同じ
	// var n = 1
	// n = 2
	// fmt.Println(1)
}

func do_something2() {
	var n = 1
	defer func() {
		fmt.Println(n)
	}()
	n = 2

	// 下記と同じ
	// var i = 1
	// i = 2
	// func() {
	// 	fmt.Println(i) // 2
	// }()
}
