package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	first := make(chan int)
	second := make(chan int)

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() {
		for _, v := range arr {
			first <- v
		}
		close(first)
	}()

	go func() {
		for i := range first {
			second <- i * 2
		}
		close(second)
	}()

	for i := range second {
		fmt.Println(i)
	}
}
