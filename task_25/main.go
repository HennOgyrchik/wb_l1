package main

import (
	"context"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {

	sleep(5 * time.Second)
}

func sleep(timeout time.Duration) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	<-ctx.Done()
}
