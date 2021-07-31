package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const StreamQty = 20

func main() {
	var counter int64 = 0

	for i := 1; i <= StreamQty; i += 1 {
		go func(i int) {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Start stream ", i)
			time.Sleep(time.Duration(i * 200) * time.Millisecond) // самый долгий поток спит 4 секунды
			atomic.AddInt64(&counter, -1)
			fmt.Println("Finish stream ", i)
		}(i)
	}

	for {
		cnt := atomic.LoadInt64(&counter)
		if cnt == 0 {
			break
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
	 fmt.Println("All finished")
}
