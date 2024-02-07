package wc_tool

import (
	"fmt"
	"os"
	"testing"
)

func TestWc(t *testing.T) {

	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err.Error())
	}
	fileStr := string(file)
	fmt.Println(fileStr)
}
