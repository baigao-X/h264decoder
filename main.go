package main

import (
	"fmt"
	"h264decoder/src/decoder"
)

func main() {
	fmt.Println("Hello, h264decoder!")

	testfile := "./tests/test.264"
	err := decoder.OpenAndParseH264(testfile)
	if err != nil {
		fmt.Printf("OpenAndParseH264 '%s': %w", testfile, err)
		return
	}

	// 打开文件
}
