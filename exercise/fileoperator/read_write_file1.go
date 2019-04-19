package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/**
将整个文件内容读到一个字符串里
*/

func main() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"

	//带缓冲的读取
	//buf := make([]byte, 1024)
	//n, err := inputReader.Read(buf);
	//if n == 0 {
	//	break
	//}

	//第一个返回值类型为 []byte
	buf, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}

	fmt.Printf("%s\n", string(buf))

	//将buf值写入文件
	//若写入文件不存在，则会创建
	err = ioutil.WriteFile(outputFile, buf, 0644)

	if err != nil {
		panic(err.Error())
	}
}
