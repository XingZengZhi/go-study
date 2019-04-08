package get_start

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	fmt.Println(Vertex{1, 2})
	// 修改Y的值
	a.Y = 10
	fmt.Println(a.X)
	fmt.Println(a.Y)

	// 结构体指针
	p := &a
	p.X = 1e9
	fmt.Println(a)
}
