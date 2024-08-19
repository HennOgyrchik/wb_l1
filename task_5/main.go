package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	ch := make(chan int)
	defer close(ch)
	// контекст завершится по истечении n секунд
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	w := worker{
		ch:  ch,
		ctx: ctx,
		wg:  &wg,
	}

	wg.Add(2)
	go w.sending()
	go w.reading()

	wg.Wait()

}

type worker struct {
	ch  chan int
	ctx context.Context
	wg  *sync.WaitGroup
}

// sending отправляет последовательные данные в канал, либо ожидает завершение контекста
func (w worker) sending() {
	defer w.wg.Done()
	defer fmt.Println("sending close")
	i := 0

	for {
		select {
		case w.ch <- i:
			i++
		case <-w.ctx.Done():
			return
		}
	}
}

// reading читает данные из канала, либо ожидает завершение контекста
func (w worker) reading() {
	defer w.wg.Done()
	defer fmt.Println("reading close")
	for {
		select {
		case tmp := <-w.ch:
			fmt.Println(tmp)
		case <-w.ctx.Done():
			return
		}
	}
}
