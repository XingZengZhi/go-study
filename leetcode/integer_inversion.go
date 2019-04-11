package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	s := strings.Split(strconv.Itoa(x), "")
	r := make([]string, len(s))
	var l = 0
	if strings.Contains(strconv.Itoa(x), "-") {
		if r[0] != "-" {
			r[0] = "-"
		}
		l = 1
	} else {
		l = 0
	}
	for i, j := len(s)-1, l; i >= l; i-- {
		r[j] = s[i]
		j++
	}
	intRever, _ := strconv.Atoi(strings.Join(r, ""))
	if -math.Pow(2, 31) >= float64(intRever) || float64(intRever) >= (math.Pow(2, 31)-1) {
		return 0
	}
	return intRever
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(
		-2147483648))
	fmt.Println(reverse(
		1534236469))
}
