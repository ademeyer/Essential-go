package main

import (
        "encoding/csv"
        "os"
        "io"
        "fmt"
        "log"
    )

/*
* Breakup of one channel into multiple ones by
* distributing each value.
*           | ----->
* --------> | ----->
*           | ----->
*/

func main() {
    err := run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("All completed, exiting")
}

func run() error {
    ch1, err := read("../fruits.csv")
    if err != nil {
        return err
    }
    
    br1 := breakup("1", ch1)
    br2 := breakup("2", ch1)
    br3 := breakup("3", ch1)
    
    for {
        if br1 == nil && br2 == nil && br3 == nil {
                break
        }
        
        select {
            case _,ok := <-br1:
                if !ok {
                    br1 = nil
                }
            case _, ok := <-br2:
                if !ok {
                    br2 = nil
                }
            case _, ok := <-br3:
                if !ok {
                    br3 = nil
                }
        }
    }
    return nil
}

func breakup(workerID string, chI <-chan []string) chan struct{} {
    ch := make(chan struct{})

    go func(){
        for v := range chI {
            fmt.Println(workerID, v)
        }
        close(ch)
    }()

    return ch
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
