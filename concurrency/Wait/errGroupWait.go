package main

import (
      "fmt"
      "errors"
      "encoding/csv"
      "io"
      "os"
      "sync"

      "golang.org/x/sync/errgroup"
    )

func main() {
  wait := errGroups() 

  <- wait
}

func errGroups() <-chan struct{}{
  ch := make(chan struct{}, 1)

  var g errgroup.Group

  for _, files := range []string{"../file1.csv", "../file2.csv", "../file3.csv"} {
    f := files
    g.Go(
        func() error {
          ch, err := read(f)
          if err != nil {
            return fmt.Errorf("erorr reading %w", err)
          }

          for line := range ch {
            fmt.Println(line)
          }

          return nil
        })
  }

  go func(){
    if err := g.Wait(); err != nil {
      fmt.Println("Error reading files %v", err)
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
