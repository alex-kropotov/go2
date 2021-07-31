package main

import (
	"fmt"
	"sync"
)

var mt sync.Mutex

func main() {
	mt.Lock()
	defer func() {
		mt.Unlock()
		fmt.Println("mt.Unlock()")
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	panic("Panic!")
}

