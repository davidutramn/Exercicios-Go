package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	nOfGoroutines := 1000
	c := 0

	wg.Add(nOfGoroutines)

	// Mutual Exclusion
	// Apenas uma goroutine tem acesso a váriavel por vez
	// evita o data race
	for i := 0; i < nOfGoroutines; i++ {
		go func() {
			runtime.Gosched()
			mu.Lock()
			c++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Contador:", c)
	fmt.Println("Esperado:", nOfGoroutines)
}
