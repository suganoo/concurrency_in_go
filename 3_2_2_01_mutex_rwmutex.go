package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetric sync.WaitGroup
	// increment
	for i := 0 ; i < 5 ; i++ {
		arithmetric.Add(1)
		go func() {
			defer arithmetric.Done()
			increment()
		}()
	}

	// decrement
	for i := 0 ; i < 5 ; i++ {
		arithmetric.Add(1)
		go func() {
			defer arithmetric.Done()
			decrement()
		}()
	}

	arithmetric.Wait()

	fmt.Println("Arithemetric complete.")
} 
