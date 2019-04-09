package main

import (
	"fmt"
	"time"
)

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300), "\n")
	computerTime()
}

func Adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

func computerTime() {
	var f = Adder()
	start := time.Now()
	fmt.Print(f(300))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
