package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	//fmt.Println("Wait...")
	//time.Sleep(2 * time.Second)
	instance := myPool.Get()
	//fmt.Println("Wait...")
	//time.Sleep(2 * time.Second)
	myPool.Put(instance)
	//fmt.Println("Wait...")
	//time.Sleep(2 * time.Second)
	myPool.Get()
}
