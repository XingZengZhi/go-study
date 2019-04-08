package get_start

import "fmt"

func main() {
	/**
	& 符号会生成一个指向其作用对象的指针（地址）
	* 符号表示指针指向的底层的值
	*/
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
