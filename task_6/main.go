package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

//Реализовать все возможные способы остановки выполнения горутины.
// Странное задание. Придумать можно все что угодно

func main() {
	var wg sync.WaitGroup

	//1) c использованием канала
	ch := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Первая горутина закрыта")

		//по необходимости можно обернуть в for или select
		<-ch
	}()

	ch <- true

	//2) с использованием контекстов. По идее почти то же, что и первый способ
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Вторая горутина закрыта")

		<-ctx.Done()
	}()

	cancel()

	// 3) самостоятельное завершение
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Третья горутина закрыта")

	}()

	// 4) Goexit завершает работу вызывающей его программы.
	//Это не затрагивает ни одну другую программу.
	//Goexit запускает все отложенные вызовы перед завершением работы программы.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Четвертая горутина закрыта")
		runtime.Goexit()

	}()

	wg.Wait()
}
