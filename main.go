package main

import (
	"fmt"
	"time"
)

func main() {
	count := 100
	var workers = make(chan int, 5)

	for i := 1; i <= count; i++ {
		workers <- i

		go func(job int){
			defer func() {
				<-workers
			}()
			fmt.Println(job)
		}(i)
	}
	time.Sleep(2 * time.Second)
}
