package main

import "fmt"

type Vertexnew struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertexnew)
	m["Bell Labs"] = Vertexnew{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
