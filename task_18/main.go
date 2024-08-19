package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое
*/

type counter struct {
	mu    *sync.Mutex
	value int
}

func (c *counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	var mu sync.Mutex
	c := counter{
		value: 0,
		mu:    &mu,
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.increment(&wg)
	}

	wg.Wait()
	fmt.Println(c.value)
}
