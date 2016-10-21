package main

type position struct {
  x, y int
}

type Queen struct {
  position position
}

func (r Queen) getPosition() position {
  return r.position
}

func (r Queen) getColumn() int {
  return r.position.x
}

func (r Queen) isDiagonal(q Queen) bool {
  return abs(r.position.x-q.position.x) == abs(r.position.y-q.position.y)
}

// math.Abs work only with float64 ??
// todo: import math/big for *int
func abs(x int) int {
  ret := x
  if ret < 0 {
    ret *= -1
  }
  return ret
}
