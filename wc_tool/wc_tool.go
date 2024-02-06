package wc_tool

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// TODO: Can use binary search to fasten it more.
func Wc() {

	bytesArg := "c"
	linesArg := "l"
	wordArg := "w"
	characterArg := "m"

	args := flag.String(bytesArg, "", "Args to the wc command")
	fileName := flag.Arg(1)
	file, _ := os.ReadFile(fileName)
	fileStr := string(file)

	fmt.Println(flag.Args())

	newLineCount := strings.Count(fileStr, "\r\n")
	fmt.Println(newLineCount)

	wordCount := len(strings.Fields(fileStr))
	fmt.Println(wordCount)

	fmt.Println(utf8.RuneCountInString(fileStr))

	switch *args {
	case bytesArg:
		fmt.Println(len(file))

	case linesArg:
		newLineCount := strings.Count(fileStr, "\r\n")
		fmt.Println(newLineCount)

	case wordArg:
		wordCount := len(strings.Fields(fileStr))
		fmt.Println(wordCount)

	case characterArg:
		fmt.Println(utf8.RuneCountInString(fileStr))

	default:
		fmt.Println(len(file))

		newLineCount := strings.Count(fileStr, "\r\n")
		fmt.Println(newLineCount)

		wordCount := len(strings.Fields(fileStr))
		fmt.Println(wordCount)

		fmt.Println(utf8.RuneCountInString(fileStr))
	}

}
