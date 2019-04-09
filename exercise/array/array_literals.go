package main

import "fmt"

func main() {
	//数组常量
	var arrayKeyValue = [5]string{3: "chirs", 4: "Ron"}
	for i := 0; i < len(arrayKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrayKeyValue[i])
	}
}

//第一种定义：
var arrAge = [5]int{18, 20, 15, 22, 16}

//第二种定义：
var arrAge2 = [...]int{5, 6, 7, 8, 22}

//第三种定义：只有3和4被赋予实际的值，其他元素都置为空的字符串
var arrKeyValue = [5]string{3: "Chirs", 4: "Ron"}
