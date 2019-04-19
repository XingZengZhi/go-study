package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
使用切片拷贝文件内容
拷贝中文有问题
*/

func main() {
	inputFile, _ := os.Open("input.dat")
	outputFile, _ := os.OpenFile("output.dat", os.O_CREATE|os.O_WRONLY, 0666)

	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}

		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputWriter.Flush()
	}

	fmt.Println("Conversion done")
}
