package main

import (
	"fmt"
	"math/rand"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	arr := []int{5, 24, 6, 7, 7, 5, 4, 4, 65, 78, 9, 0, 8, 45, 34563, 42, 7, 34, 890, 5, 34, 3}
	fmt.Println(quicksort(arr))
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
