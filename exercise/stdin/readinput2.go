package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input2      string
	err         error
)

func main() {
	//创建读取器，与标准输入绑定
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	//从输入中读取内容，直到碰到指定的字符
	//如果读取过程中没有碰到指定的字符串，将返回错误 err != nil
	input2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input2)
	}
}
