package wc_tool

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// Wc is a function that performs word count operations on a file.
// It takes command line arguments to specify the file and the operation to be performed.
// The supported options are:
//   -c: Count the number of bytes in the file.
//   -l: Count the number of lines in the file.
//   -w: Count the number of words in the file.
//   -m: Count the number of characters in the file.
// If no option is provided, it counts the number of bytes, lines, words, and characters in the file.

func Wc() {

	noOfBytes := flag.Bool("c", false, "number of bytes in a file")
	noOfLines := flag.Bool("l", false, "number of bytes in a file")
	noOfWords := flag.Bool("w", false, "number of bytes in a file")
	noOfCharacter := flag.Bool("m", false, "number of bytes in a file")

	flag.Parse()

	var filePath string
	var fileStr string

	args := flag.Args()

	if len(args) > 0 {
		filePath = args[0]
		_, fileStr = ReadFIle(filePath)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err.Error())
		}

		fileStr = string(stdin)
	}

	switch {
	case *noOfBytes:
		fmt.Println(CalculateNoOfBytes(fileStr))

	case *noOfLines:
		fmt.Println(CalculateNoOfLines(fileStr))

	case *noOfWords:
		fmt.Println(CalculateNoOfWords(fileStr))

	case *noOfCharacter:
		fmt.Println(CalculateNoOfCharacter(fileStr))

	default:
		fmt.Println(CalculateNoOfBytes(fileStr))

		fmt.Println(CalculateNoOfLines(fileStr))

		fmt.Println(CalculateNoOfWords(fileStr))

		fmt.Println(CalculateNoOfCharacter(fileStr))
	}

	if len(args) > 0 {
		fmt.Println(args[0])
	}

}

func CalculateNoOfCharacter(filestr string) int {
	return utf8.RuneCountInString(filestr)
}

func CalculateNoOfWords(fileStr string) int {
	wordCount := len(strings.Fields(fileStr))
	return wordCount
}

func CalculateNoOfBytes(fileStr string) int {
	return len(fileStr)
}

func CalculateNoOfLines(filestr string) int {
	newLineCount := strings.Count(filestr, "\r\n")
	return newLineCount

}

func ReadFIle(fileName string) ([]byte, string) {

	var fileStr string

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	fileStr = string(file)
	return file, fileStr
}
