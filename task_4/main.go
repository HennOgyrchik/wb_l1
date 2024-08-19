package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	// канал
	mainCh := make(chan int)
	defer close(mainCh)

	var wg sync.WaitGroup

	// Для закрытия worker'ов используется NotifyContext
	// NotifyContext возвращает копию родительского контекста,
	//которая помечается как выполненная (ее канал Done закрыт) при поступлении одного из перечисленных сигналов,
	//при вызове возвращаемой функции stop или при закрытии канала Done родительского контекста,
	//в зависимости от того, что произойдет раньше.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	w := worker{
		ch:  mainCh,
		ctx: ctx,
		wg:  &wg,
	}
	w.run(3)

	// постоянная запись в канал (главный поток по ТЗ), либо ожидание завершения контекста
	for {
		select {
		case mainCh <- rand.Int():
		case <-ctx.Done():
			//ожидаем завершение всех потоков
			wg.Wait()
			return
		}
	}
}

// worker структура
type worker struct {
	ch  chan int
	ctx context.Context
	wg  *sync.WaitGroup
}

// Run запускает count горутин
func (w worker) run(count int) {
	for i := 0; i < count; i++ {
		w.wg.Add(1)
		go w.reading()
	}
}

// reading пытается бесконечно считывать из канала, либо ожидает завершение контекста
func (w worker) reading() {
	defer w.wg.Done()
	defer fmt.Println("reading close")

	for {
		select {
		case data := <-w.ch:
			fmt.Println(data)
		case <-w.ctx.Done():
			return
		}
	}
}
