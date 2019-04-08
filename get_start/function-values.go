package get_start

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	}

	fmt.Println(hypot(3, 4))
}
