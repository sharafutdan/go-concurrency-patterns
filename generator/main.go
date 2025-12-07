package main

import (
	"fmt"
	"time"
)

func generatorFunc(message string) <-chan string {
	resChan := make(chan string)
	go func() {
		for i := 0; ; i++ {
			resChan <- fmt.Sprintf("%s, %d", message, i)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return resChan
}

func main() {
	joe := generatorFunc("joe")
	mary := generatorFunc("mary")
	i := 0
	for i < 10 {
		fmt.Printf("message received: %v\n", <-joe)
		fmt.Printf("message received: %v\n", <-mary)
		i++
	}
	fmt.Println("exiting")
}
