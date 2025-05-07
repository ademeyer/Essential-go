package main

import(
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Wait for two go-routines
	
	// First go-routine
	go func(){
		count("sheep", 10)
		wg.Done() // mark routine as done
	}()
	
	// Second go-routine
	go func() {
		count("fish", 100)
		wg.Done() // mark routine as done
	}()

	wg.Wait()
}

func count(name string, sleepTime time.Duration){
	for i := 0; i <= 20; i++{
		fmt.Println(i, name)
		time.Sleep(time.Millisecond * sleepTime)
	}
}
