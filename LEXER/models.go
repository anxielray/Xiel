package LEXER

import (
	"strings"
)

type LEXER struct {
	Source_code  string
	Split_code   [][]string
	Final_tokens [][]string
	Key_words    KEY_WORDS
}

type KEY_WORDS struct {
	Function_key      string
	Function_operands string
	Indentation       string
	Dentation         string
	Identifier        string
}

var sub_split [][]string

func Sub_split(split []string) [][]string {

	// if strings.Contains(split[0], "print"){
	// 	split = append([]string{"IDENTIFIER"}, split...)
	// }
	for i := 0; i < len(split[0]); i++ {
		s := split[0][i]
		tmp := []string{}
		if s == '(' {
			tmp = append(tmp, split[0][:i])
			sub_split = append([][]string{tmp}, sub_split...)
			// split = append([]string{"OPENING_PARENTHESIS"}, split...)
		} else if s == ')' {
			tmp = append(tmp, split[0][:i])
			sub_split = append([][]string{tmp}, sub_split...)
			// split = append([]string{"CLOSING_PARENTHESIS"}, split...)
		} else if s == '"' {
			tmp = append(tmp, split[0][:i])
			sub_split = append([][]string{tmp}, sub_split...)
			// split = append([]string{"QUOTE"}, split...)
		} else {
			continue
		}
		tmp = nil
		// i = 0
		split[0] = split[0][i:]

	}
	return sub_split
}

func Label(split_code [][]string) [][]string {
	//if the keyword that is found is anx, label as keyword and the next as identifier
	var new_words [][]string
	for _, w := range split_code {
		//other checks before you conclude with the identifier section

		if w[0] != "IDENTIFIER" && w[0] != ".>" && w[0] != "anx" && w[0] != "if" && w[0] != "elif" && w[0] != "else" {
			w = append([]string{"IDENTIFIER"}, w...)
		}
		new_words = append(new_words, w)
	}
	return new_words
}

type Methods_on_lexer interface {
	Split_white_space()
}

func (l *LEXER) Split_white_space() {
	code_split := strings.SplitSeq(l.Source_code, "\n")
	for line := range code_split {
		words := strings.Fields(line)
		for _, w := range words {
			l.Split_code = append(l.Split_code, Sub_split([]string{w})...)
		}
	}
	l.Split_code = Label(l.Split_code)
}
