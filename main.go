package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type TimedError struct {
	errTime time.Time
}

func (e *TimedError) Error() string {
	return fmt.Sprintf("Неожиданная ошибка, " + e.errTime.String())
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			e := &TimedError{
				errTime: time.Now(),
			}
			log.Fatalln(e)
		}
	}()
	createFile("test.txt")
	WrongDivide()

}

func WrongDivide() {
	for i := 10; i >= 0; i-- {
		fmt.Println(7/i)
	}
}

func createFile(name string) {
	file, _ := os.Create(name)
	defer func() {
		file.Close()
		log.Println("Закрыли файл")
	}()
}