package main

import (
	"fmt"
	"slices"
)

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(set(arr))
}

func set(arr []string) []string {
	var result []string
	for _, v := range arr {
		if i := slices.Index(result, v); i == -1 {
			result = append(result, v)
		}
	}
	return result
}
