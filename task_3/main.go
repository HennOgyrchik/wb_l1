package main

import (
	"fmt"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	// канал для передачи данных между горутинами
	ch := make(chan int)
	defer close(ch)

	//цикл по массиву
	for _, v := range arr {
		//отправка результата выполнения square() в канал отдельными горутинами
		go func() {
			ch <- square(v)
		}()
	}

	// итоговый подсчет
	result := 0
	for i := 0; i < len(arr); i++ {
		tmp := <-ch
		result += tmp
	}

	fmt.Println(result)
}

func square(i int) int {
	return i * i
}
