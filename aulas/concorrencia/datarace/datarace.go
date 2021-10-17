package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nOfGoroutines := 1000
	c := 0 // armazena a quantidade de vezes que as goroutines executaram

	wg.Add(nOfGoroutines)

	// Data Race
	// acontece quando 2 goroutines acessam a mesma variavel concorrentemente
	// e uma delas esta alterando o valor
	for i := 0; i < nOfGoroutines; i++ {
		// Para testar se acontece o data race execute a aplicação com o comando
		// go run -race datarace.go
		go func() {
			runtime.Gosched()
			c++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Contador:", c)
	fmt.Println("Esperado:", nOfGoroutines)
}
