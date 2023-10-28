package main

import "fmt"

func main() {
finished:
	// ラベル付きループ
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if true {
				fmt.Println("break!!")
				break finished // 内側のループから外側のループに対して break 文を実行
			}
		}
	}
}
