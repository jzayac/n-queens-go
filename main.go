package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
  "time"
)

var nQueen int
var solution int

func intInSlice(i int, list []int) bool {
  for _, b := range list {
    if b == i {
      return true
    }
  }
  return false
}

func findNext(queens []Queen, column []int) {
  if len(queens) == nQueen {
    _ = "breakpoint"
    solution++
  } else {
    for i := 0; i < nQueen; i++ {
      row := len(queens)
      current := Queen{position{i, row}}
      if !conflict(queens, current, column) {
        findNext(append(queens, current), append(column, i))
      }
    }
  }
}

func conflict(queens []Queen, current Queen, column []int) bool {
  if intInSlice(current.getColumn(), column) {
    return true
  }
  for _, q := range queens {
    if current.isDiagonal(q) {
      return true
    }
  }
  return false
}

func start() {
  for i := 0; i < int(math.Ceil(float64(nQueen)/2.0)); i++ {
    column := []int{i}
    queens := make([]Queen, 1, 12)
    queens[0] = Queen{position{i, 0}}
    findNext(queens, column)

    if i != 0 && int(math.Floor(float64(nQueen)/2.0)) == i+1 {
      solution *= 2
    }
  }
}

func main() {
  arg, err := strconv.Atoi(os.Args[len(os.Args)-1:][0])
  if err != nil {
    fmt.Println("last arg must be integer")
    os.Exit(2)
  }
  nQueen = arg
  solution = 0
  startTime := time.Now().UnixNano() / int64(time.Millisecond)

  fmt.Printf("find: %d\n", nQueen)

  start()
  fmt.Printf("time: %d\n", (time.Now().UnixNano()/int64(time.Millisecond) - startTime))
  fmt.Printf("-----------------\n")
  fmt.Printf("result: %d\n", solution)
  fmt.Printf("-----------------\n")
}
