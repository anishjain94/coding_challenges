`wc_tool` is a Go package that performs word count operations on a file.
It provides functionality to count the number of bytes, lines, words, and characters in a file.

## Usage

To use `wc_tool`, you need to provide command line arguments to specify the file and the operation to be performed. The supported options are:

- `-c`: Count the number of bytes in the file.
- `-l`: Count the number of lines in the file.
- `-w`: Count the number of words in the file.
- `-m`: Count the number of characters in the file.

If no option is provided, it counts the number of bytes, lines, words, and characters in the file.

## Build Command

`go build main.go`

## Run

`./main test.txt`
