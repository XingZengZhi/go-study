package get_start

import "fmt"

func main() {
	var a []int
	fmt.Println("a", a)

	a = append(a, 0, 9)
	fmt.Println("a", a)
}
