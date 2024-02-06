package wc_tool

import (
	"fmt"
	"os"
	"strings"
)

func Wc() {


	
	file, _ := os.ReadFile("test.txt")
	fmt.Println(len(file))
	fileStr := string(file)

	newLineCount := strings.Count(fileStr, "\r\n")
	fmt.Println(newLineCount)

	wordCount := len(strings.Fields(fileStr))
	fmt.Println(wordCount)

}
