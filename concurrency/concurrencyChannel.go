package main

import(
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count(c)	

	for msg := range c { 
		fmt.Println(msg)
	}
}

func count(c chan string) {
	name := []string{"sheep", "fish", "meat", "cow", "turkey"}
	for i := 0; i < len(name); i++ {
		c <- name[i]
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
