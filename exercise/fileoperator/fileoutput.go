package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
文件写入数据
只能增加，不能减少
顶部开始写入
*/
func main() {
	//以只写模式打开文件，如果不存在则创建，第三个参数为文件权限
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	//创建写入器（缓冲区）
	outputWriter := bufio.NewWriter(outputFile)

	outputString := "hello world111!\n"

	for i := 0; i < 1; i++ {
		//写入字符串
		//outputFile.WriteString("hello, world in a file\n")
		//fmt.Fprintf(outputFile, "Some test data.\n")
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
