package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

/**
读取压缩文件
读到了压缩文件中gzipped.txt的内容，
但是没读取到首行和最后一行的数据
*/
func main() {
	fName := "gzipped.tar.gz"
	var r *bufio.Reader

	fi, err := os.Open(fName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)

	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
