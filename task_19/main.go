package main

import (
	"fmt"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	str := "главрыба"

	fmt.Println(reverse(str))
}

func reverse(str string) string {
	arr := []rune(str)

	var res []rune

	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return string(res)
}
