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

func FanIn(channels ...<-chan string) <-chan string {
	c := make(chan string)
	for _, channel := range channels {
		go func() {
			for {
				c <- <-channel
			}
		}()
	}
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
