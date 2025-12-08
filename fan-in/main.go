package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatorFunc(message string) <-chan string {
	resChan := make(chan string)
	go func() {
		for i := 0; ; i++ {
			resChan <- fmt.Sprintf("%s, %d", message, i)
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		}
	}()
	return resChan
}

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		select {
		case v := <-input1:
			c <- v
		case v := <-input2:
			c <- v
		}
	}()
	return c
}

func main() {
	joe := generatorFunc("joe")
	mary := generatorFunc("mary")
	for i := 0; i < 10; i++ {
		fmt.Println(<-FanIn(joe, mary))
	}
	fmt.Println("exiting")
}
