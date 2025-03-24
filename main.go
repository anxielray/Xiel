package main

import (
	"fmt"
	"os"
	p "xiel/LEXER"
)

func main() {
	var p p.LEXER
	p.Source_code = read_file(os.Args[1])
	p.Split_white_space()
	fmt.Println(p.Split_code)
}

func read_file(file_path string) string {
	content, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("[ERROR]: could not open file")
		os.Exit(1)
	}
	return string(content)
}
