package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a := 1
	b := 3

	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
