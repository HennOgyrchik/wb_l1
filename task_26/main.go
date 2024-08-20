package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	s := []string{"abcd", "abCdefAaf", "aabcd", "dAr3b"}

	for _, str := range s {
		fmt.Printf("%s - %t\n", str, uniq(str))
	}
}

func uniq(str string) bool {
	s := strings.ToLower(str)

	for _, r := range s {
		if count := strings.Count(s, string(r)); count > 1 {
			return false
		}
	}

	return true
}
