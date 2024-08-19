package main

import "sync"

//Реализовать конкурентную запись данных в map.

type repository struct {
	mu   *sync.Mutex
	data map[int]int
	wg   *sync.WaitGroup
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	rep := repository{
		data: make(map[int]int),
		mu:   &mu,
		wg:   &wg,
	}

	wg.Add(3)
	go rep.writer()
	go rep.writer()
	go rep.writer()

	wg.Wait()
}

func (r *repository) writer() {
	defer r.wg.Done()
	for i := 0; i < 50; i++ {
		//Для блокирования доступа к общему разделяемому ресурсу
		//у мьютекса вызывается метод Lock(), а для разблокировки доступа - метод Unlock().
		r.mu.Lock()
		r.data[i] = i
		r.mu.Unlock()
	}
}
