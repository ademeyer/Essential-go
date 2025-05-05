package main

import (
    "fmt"
    "encoding/csv"
    "io"
    "os"
    "log"
    "sync"
    )

/*
* FanIn: Consolidation of multiple channels into 
* one channel by multiplexing each received value
*
* ---> |
* ---> | ----->
* ---> |
*/

func main() {
    err := run()
    if err != nil {
        log.Fatal(err)
    }
}

func run() error {
    ch1, err := read("../fruits.csv")
    if err != nil {
        return err
    }

    ch2, err := read("../fruits.csv")
    if err != nil {
        return err
    }

    // Create graceful shutdown watch
    exit := make(chan struct{})
    ch := merge(ch1, ch2)

    go func(){
        for v := range ch {
            fmt.Println(v)
        }

        close(exit)
    }()
    <- exit
    return nil
}


func merge(cs ...<-chan []string) <-chan []string {
    var wg sync.WaitGroup
    
    // channel to collect all other channels
    out := make(chan []string)
    
    // this function will called by goroutine
    send := func(c <-chan []string){
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    
    // Create the wait number of cs passed into merge
    wg.Add(len(cs))
    
    // Collect all cs channels and merge into out channel
    // in a goroutine
    for _, c := range cs {
        go send(c)
    }

    // Spine another goroutine to wait for 
    // send to merge all cs channel 
    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

func read(file string) (<-chan []string, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }

    ch := make(chan []string)
    
    cr := csv.NewReader(f)

    go func() {
        for {
            record, err := cr.Read()
            if err == io.EOF {
                close(ch)
                return
            }

            ch <- record
        }
    }()

    return ch, nil
}
