package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	x := 5

	fmt.Println(find(arr, x))
}

func find(arr []int, target int) int {
	if target > arr[len(arr)-1] {
		return -1
	}

	mid := len(arr) / 2
	switch {
	case target > arr[mid]:
		return mid + find(arr[mid:], target)
	case target < arr[mid]:
		return find(arr[:mid], target)
	default:
		return mid
	}
}
