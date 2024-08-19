package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	x := int64(54)
	i := 1

	fmt.Printf("%064b\n", x)

	//инвертирует i-й бит
	x ^= 1 << i

	fmt.Printf("%064b\n", x)
}
