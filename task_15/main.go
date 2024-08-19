package main

import (
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

var justString string

func someFunc() {
	//неаккуратное обращение с побитовым сдвигом может привести к переполнению или утечке памяти.
	//тут создается строка длинной 1024, при использовании первых 100 символов
	v := createHugeString(1 << 10)
	//в случае, если v будет короче 100, случится паника (выход за пределы)
	justString = v[:100]

	justString = createHugeString(100)

}

func createHugeString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	someFunc()
}
