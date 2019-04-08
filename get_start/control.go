package get_start

import (
	"fmt"
	"runtime"
)

func main() {
	i, count := 0, 0

	for ; i <= 100; i++ {
		count += i
	}

	defer fmt.Println(count)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s", os)
	}
}
