package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//以只读模式打开文件
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	//循环读取文件内容
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		//io.EOF的值是 true
		if readerError == io.EOF {
			return
		}
	}
}
