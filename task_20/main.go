package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	str := "snow dog sun"

	fmt.Println(reverse(str))
}

func reverse(str string) string {
	arr := strings.Fields(str)

	var res []string

	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}

	return strings.Join(res, " ")
}
