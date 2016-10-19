package main

import (
	"fmt"
	"math"
	"time"
)

type position struct {
  x, y int
}

type queen struct {
  position position
}

func (r queen) getPosition() position {
  return r.position
}

func (r queen) getColumn() int {
  return r.position.x
}

func (r queen) isDiagonal(q queen) bool {
  return abs(r.position.x - q.position.x) == abs(r.position.y - q.position.y)
}

// math.Abs work only with float64 ??
// todo: import math/big for *int
func abs(x int) int {
  ret := x;
  if (ret < 0) {
	  ret *= -1
	}
  return ret
}

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

func findNext(queens []queen, column []int) {
  if len(queens) == nQueen {
	  _ = "breakpoint"
	  solution++ 
	} else {
	  for i := 0; i < nQueen; i++ {
		  row := len(queens)
      current := queen{position{i, row}}
			if (!conflict(queens, current, column)) {
				findNext(append(queens, current), append(column, i))
			}
		}
	}
}

func conflict(queens []queen, current queen, column []int) bool {
  if (intInSlice(current.getColumn(), column)) {
	  return true
	}
	for _, q := range queens {
    if (current.isDiagonal(q)) {
		  return true
		}
	}
	return false
}


func start() {
	for i := 0; i < int(math.Ceil(float64(nQueen)/2.0)); i++ {
    column := []int{i}
		queens := make([]queen, 1, 12)
    queens[0] = queen{position{i, 0}}
		findNext(queens, column)

    if (i != 0 && int(math.Floor(float64(nQueen)/2.0)) == i + 1) {
		  solution *= 2
		}
	}
}

func main() {
	nQueen = 15
	solution = 0
  startTime := time.Now().UnixNano() / int64(time.Millisecond)

  fmt.Printf("find: %d\n", nQueen)

	start()
  fmt.Printf("time: %d\n", (time.Now().UnixNano() / int64(time.Millisecond)  - startTime))
	fmt.Printf("-----------------\n")
	fmt.Printf("result: %d\n", solution)
	fmt.Printf("-----------------\n")
}
