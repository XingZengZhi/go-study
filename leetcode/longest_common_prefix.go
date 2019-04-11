package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := []string{"flower","flow","flight"}
	//s := []string{"aa","ab"}
	s := []string{"aca", "cba"}
	//s := []string{}
	//s := []string{"a"}
	//s := []string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	y, index, nums := len(strs[0]), 0, 0
	s := ""
	for i, v := range strs {
		if y > len(v) {
			y = len(v)
			index = i
		}
	}
	for i, v := range strings.Split(strs[index], "") {
		nums = 0
		for j, k := range strs {
			if index == j {
				continue
			}
			for o, n := range strings.Split(k, "") {
				if i != o && v == n {
					continue
				}
				if v == n {
					nums++
					if nums == len(strs)-1 {
						s += n
					}
					continue
				}
			}

		}
		if nums == 0 {
			break
		}
	}

	return s
}

//func longestCommonPrefix(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//	i := 0
//	for ;; {
//		for j := 0; j < len(strs); j++ {
//			if i >= len(strs[j]) {
//				return strs[0][:i]
//			}
//			if strs[j][i] != strs[0][i] {
//				return strs[0][0:i]
//			}
//		}
//		i++
//	}
//	return strs[0]
//}
