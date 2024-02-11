package main

import (
	"coding_challenges/wc_tool"
	"testing"
)

func TestGetFileStats(t *testing.T) {

	_, fileStr := wc_tool.ReadFIle("test.txt")

	wantedBytes := 342190

	wantedLines := 7145

	wantedWords := 58164

	wantedChar := 339292

	if wantedBytes != wc_tool.CalculateNoOfBytes(fileStr) {
		t.Errorf("got %+v want %+v", wc_tool.CalculateNoOfBytes(fileStr), wantedBytes)
	}

	if wantedLines != wc_tool.CalculateNoOfLines(fileStr) {
		t.Errorf("got %+v want %+v", wc_tool.CalculateNoOfLines(fileStr), wantedBytes)
	}

	if wantedWords != wc_tool.CalculateNoOfWords(fileStr) {
		t.Errorf("got %+v want %+v", wc_tool.CalculateNoOfWords(fileStr), wantedBytes)
	}

	if wantedChar != wc_tool.CalculateNoOfCharacter(fileStr) {
		t.Errorf("got %+v want %+v", wc_tool.CalculateNoOfCharacter(fileStr), wantedBytes)
	}

}
