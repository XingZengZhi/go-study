package _var

import "os"

//var a int
//var b bool
//var str string

//or
//因式分解关键字写法一般用于声明全局变量
//当变量被声明过后，系统自动赋予它该类型的零值：int为0，float为0.0，bool为false，
//string为空字符串，指针为nil。所有的内存在go中都是经过初始化的。
//变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词首字母大写。
//如果希望全局变量能够被外部包所使用，则需要将首个单词的首字母也大写。
var (
	a   int
	b   bool
	str string
)

// 这种写法主要用于声明包级别的全局变量，若在函数体内声明全局变量时，可以使用
//短声明方法 := ，
var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

//a := 1
