package main

import (
	"fmt"
)

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// do something more
	fmt.Println("Done.")
}
