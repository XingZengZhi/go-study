package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 0
	fmt.Println(isPalindrome(11))
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	a := strings.Split(s, "")
	if len(s) >= 1 && !strings.Contains(s, "-") {
		for i, j := 0, len(s)-1; i != j; i++ {
			if a[i] != a[j] {
				return false
			} else if i == len(s)-1 && j == 0 {
				return true
			} else {
				j--
			}
		}
		return true
	}
	return false
}

//func isPalindrome(x int) bool {
//	var s int
//	for X := x; X > 0; X /= 10 {
//		y := X % 10
//		s = s*10 + y
//	}
//	return (s == x)
//}
