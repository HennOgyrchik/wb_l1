package main

import (
	"fmt"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var x interface{}

	x = 2
	fmt.Printf("Тип переменной x: %T\n", x)

	x = "asdasd"
	fmt.Printf("Тип переменной x: %T\n", x)

	x = make(chan int)
	fmt.Printf("Тип переменной x: %T\n", x)
}
