package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		for {
			c1 <- "Hello!"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func(){
		for {
			c2 <- "Hi!"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg := <- c1:
			fmt.Printf("%v, c1: %s\n", time.Now(), msg)
                case msg := <- c2:
                        fmt.Printf("%v, c2: %s\n", time.Now(), msg)
		}
	}
}
