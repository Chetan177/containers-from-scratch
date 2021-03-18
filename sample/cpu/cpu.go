package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	t := runtime.GOMAXPROCS(0)

	fmt.Printf("GOMAXPROCS is %d\n", t)

	for i := 0; i < t; i++ {
		wg.Add(1)
		go eatcpu(&wg)
	}

	wg.Wait()
}

func eatcpu(wg *sync.WaitGroup) {
	ch := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-ch.C:
			wg.Done()
			return
		default:
		}
	}
}
