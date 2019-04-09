package main

import "fmt"

//大数组传递给函数很消耗内存，两种方法避免这种现象
//1.传递数组的指针
//2.使用数组的切片
func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array)
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
