package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	// цикл по массиву
	for _, v := range arr {
		wg.Add(1)

		//запуск функции подсчета квадрата в горутине
		go square(&wg, v)
	}

	//ожидаем завершение горутин
	wg.Wait()
}

func square(wg *sync.WaitGroup, i int) {
	fmt.Printf("%d^2 = %d\n", i, i*i)
	wg.Done()
}
