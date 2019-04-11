package main

import (
	"fmt"
)

func main() {
	a := []int{3, 3}
	//a := []int{3,2,4}
	//a := []int{2,7,11,15}
	//a := []int{-1,-2,-3,-4,-5}
	//a := []int{-3,4,3,90}
	fmt.Println(twoSum(a, 6))
}

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		c := target - v
		for j, u := range nums {
			if i == j {
				continue
			}
			if c == u {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//func twoSum(nums []int, target int) []int {
//	//b := make([]string, len(nums))
//	//p := new([]int)
//	m := make(map[int]int)
//	for i, v := range nums {
//		n := target - v
//		fmt.Println(m[n])
//		if m[n] >= 0 {
//			return []int{m[n], i}
//		}
//		m[v] = i
//	}
//	return []int{}
//}
