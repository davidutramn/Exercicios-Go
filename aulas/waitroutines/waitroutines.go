package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("First go routine")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Second go routine")
		wg.Done()
	}()

	wg.Wait()
}
