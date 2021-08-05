package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {
	const count = 10
	//var mt sync.Mutex
	var counter int64 = 0
	var v float64
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()


	for i := 1; i <= 10; i += 1 {
		go func(i int) {
			atomic.AddInt64(&counter, 1)
			time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second) // ждем от 1 до 3 секунд, потом пытаемся блокировать
			//
			// Если убрать мьютекс, вероятно будет гонка, потому что будут попытки одновременого доступа к переменой v
			//
			//mt.Lock()
			fmt.Println(i, "mt.Lock()")

			for j := 1; j < 1000000; j += 1 {
				v = float64(j) * math.Pow(float64(j), 0.5)
			}
			fmt.Println(v)
			//time.Sleep(time.Duration(rand.Intn(1)+1) * time.Second) // как будто 1-2 секунды что-то делаем
			//mt.Unlock()
			fmt.Println(i, "mt.Unlock()")
			atomic.AddInt64(&counter, -1)
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
}
