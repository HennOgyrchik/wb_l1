package main

import (
	"fmt"
	"slices"
)

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	first := []int{1, 4, 6, 3, 7, 2, 765, 10, 14}
	second := []int{9, 5, 3, 6, 8, 765, 34, 68}

	if len(first) < len(second) {
		first, second = second, first
	}

	var result []int

	for _, v := range first {
		if i := slices.Index(second, v); i != -1 {
			result = append(result, v)
		}
	}
	fmt.Println(result)
}
