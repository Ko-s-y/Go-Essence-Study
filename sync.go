package main

import (
	"fmt"
	"sync"
)

func main() {
	ParallelProcessing()
	LoopTest()
}

func ParallelProcessing() {
	var wg sync.WaitGroup
	wg.Add(1) // リファレンスカウンタを+1
	go func() {
		defer wg.Done() // リファレンスカウンタを-1
		// 重たい処理を行う
	}()

	// 別の重たい処理を行う
	wg.Wait() // リファレンスカウンタが０になるまで待つ
}

// func LoopTest() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println(i)
// 		}()
// 	}
// 	wg.Wait()
// }

// 上記 LoogTest()を修正
func LoopTest() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()
}
