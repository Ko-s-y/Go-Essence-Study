package main

import (
	"fmt"
	"sort"
)

func main() {
	set_map()
	initialize_map()
	map_enumeration()
	sort_map_enumeration()
	map_value_check()
}

func set_map() {
	map_sample := make(map[string]int, 10)
	map_sample["John"] = 21
	map_sample["Bob"] = 18
	map_sample["Mark"] = 33

	fmt.Println(map_sample)
}

func initialize_map() {
	map_sample := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}

	fmt.Println(map_sample)
	delete(map_sample, "Bob")
	fmt.Println(map_sample)
}

func map_enumeration() {
	map_sample := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}

	for k, v := range map_sample {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func sort_map_enumeration() {
	map_sample := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}

	keys := []string{}
	for k := range map_sample {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v\n", k, map_sample[k])
	}
}

func map_value_check() {
	map_sample := map[string]int{"Bob": 18}

	v, ok := map_sample["Bob"]
	if ok {
		// Bobのkeyが存在すればvalueが返される
		fmt.Println(v)
	}

	map_1 := map[string]int{"John": 18}
	fmt.Println(map_1["Who"]) // 存在しないキーを指定した場合はint型のデフォルト値０が返される

	map_2 := map[string]string{"foo": "bar"}
	fmt.Println(map_2["Who"]) // 存在しないキーを指定した場合はint型のデフォルト値の空文字が返される
}
