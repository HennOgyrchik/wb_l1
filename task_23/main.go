package main

import (
	"fmt"
	"slices"
)

/*
Удалить i-ый элемент из слайса
*/

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr, len(arr), cap(arr))

	i := 3
	result := slices.Delete(arr, i, i+1)

	fmt.Println(result, len(result), cap(result))
}
