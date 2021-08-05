package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	begin := time.Now()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go lazy(i)
	}

	wg.Wait()
	fmt.Println(time.Since(begin))
}

func lazy(i int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(i)
	wg.Done()
}
