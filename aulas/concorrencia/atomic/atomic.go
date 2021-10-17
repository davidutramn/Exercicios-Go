package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	nOfGoroutines := 1000
	var c int32

	wg.Add(nOfGoroutines)

	// Atomic
	// implementação low-level para sincronização de algoritmos
	// Recomendado apenas para casos muitos especificos
	for i := 0; i < nOfGoroutines; i++ {
		go func() {
			atomic.AddInt32(&c, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Contador:", c)
	fmt.Println("Esperado:", nOfGoroutines)
}
