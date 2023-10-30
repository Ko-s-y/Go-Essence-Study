package main

import "fmt"

// go run array.go
func main() {
	array_test()
	slice_test()
	array_slice_2d_test()
	slice_append_test()
	slice_specify_len_and_cap()
	slice_partial_reference()
	remove_slice_elements_pattern1()
	remove_slice_elements_pattern2()
	remove_slice_elements_pattern3()
}

func array_test() {
	var arr [4]int
	arr[0] = 1
	// arr[4] = 1
	// => ./array.go:13:6: invalid argument: index 4 out of bounds [0:4]

	fmt.Println(arr) // => [1 0 0 0]

	init_arr := []int{1, 3, 4, 5, 6}
	fmt.Println(init_arr)
}

func slice_test() {
	// var slc []int
	// slc[1] = 2
	// fmt.Println(slc)
	// => panic: runtime error: index out of range [1] with length 0

	slc := make([]int, 3)

	slc[0] = 1
	slc[1] = 2

	fmt.Println(slc) // => [1 2 0]
}

func array_slice_2d_test() {
	arr1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	slc2 := [][]int{
		{4, 5, 6},
		{3, 2, 1},
	}

	fmt.Println(arr1)
	fmt.Println(slc2)
}

func slice_append_test() {
	a := []int{1, 2, 3}
	fmt.Println(a)

	a = append(a, 4, 5, 6)
	fmt.Println(a)
	fmt.Printf("aの長さは %d\n", len(a))
	fmt.Println(len("aaaaaaaaaa"))
}

func slice_specify_len_and_cap() {
	a := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(a)
	}
}

func slice_partial_reference() {
	a := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(a[0:6])
	fmt.Println(a[2:5])
}

// 新しく同じかたのスライスを用意して詰め直す
// あらかじめcapを指定しておくことで無駄なアロケーションを防ぐ
func remove_slice_elements_pattern1() {
	a := []int{1, 2, 3, 4, 5, 6}
	a2 := make([]int, 0, len(a)/2)

	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			a2 = append(a2, a[i]) // 偶数のみa2に代入
		}
	}
	a = a2
}

// appendを使う
func remove_slice_elements_pattern2() {
	a := []int{1, 2, 3, 4, 5, 6}
	n := 1
	a = append(a[:n], a[n+1:]...)
	fmt.Println(a)
}

// 部分参照とcopyを使う
func remove_slice_elements_pattern3() {
	a := []int{1, 2, 3, 4, 5, 6}
	n := 1
	a = a[:n+copy(a[n:], a[n+1:])]
	fmt.Println(a)
}
