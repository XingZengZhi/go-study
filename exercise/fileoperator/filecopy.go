package main

import (
	"fmt"
	"io"
	"os"
)

/**
没成功...
*/
func main() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if src != nil {
		return
	}
	defer src.Close()

	dst, err := os.Open(dstName)
	if dst != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
