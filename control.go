package main

import "fmt"

func main() {
	i, count := 0, 0

	for ;i <= 100;i++ {
		count += i
	}

	fmt.Println(count)
}
