package main

import (
	"fmt"
	"math"
)

type Vertey struct {
	x, y float64
}

func (s *Vertey) Abs() float64 {
	return math.Sqrt(math.Pow(s.x, 2) + math.Pow(s.y, 2))
}

func main() {
	v := &Vertey{3, 4}
	fmt.Print(v.Abs())
}
