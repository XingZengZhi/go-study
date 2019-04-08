package get_start

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

const Truth = true

func main() {

	fmt.Println(add(4, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println("Go rules", Truth)

}
