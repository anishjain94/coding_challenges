package wc_tool

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func Wc() {

	args := os.Args[1:]

	fileName := "test.txt"
	option := ""

	if len(args) == 2 {
		option = args[0]
		fileName = args[1]
	} else if len(args) == 1 {
		fileName = args[0]
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	fileStr := string(file)

	switch option {
	case "-c":
		fmt.Println(len(file))

	case "-l":
		newLineCount := strings.Count(fileStr, "\r\n")
		fmt.Println(newLineCount)

	case "-w":
		wordCount := len(strings.Fields(fileStr))
		fmt.Println(wordCount)

	case "-m":
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
