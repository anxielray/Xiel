package LEXER

import (
	"slices"
	"strings"
)

type LEXER struct {
	Source_code  string
	Split_code   [][]string
	Final_tokens [][]string
}

func Sub_split(split []string) []string {
	Parenthesis := []string{"(", ")", "[", "]", "{", "}"}
	// quotes := []string{`'`, `"`}
	var sub_split, tmp []string
	//{anx, main(), .>}
	for _, s := range split {
		if contains(Parenthesis, s) {
			tmp = append(tmp, separateParenthesis(s)...)
		}
		if len(tmp) == 0 {
			tmp = append(tmp, s)
		}
		// tmp = append(tmp, separateQuotes()...)
		sub_split = append(sub_split, tmp...)
		tmp = []string{}
	}
	return sub_split
}

func separateQuotes(a string) []string {
	var result []string
	start := 0
	for i, c := range a {
		if string(c) == `'` || string(c) == `"` {
			if i > start {
				result = append(result, a[start:i])
			}
			result = append(result, string(c))
			start = i + 1
		}
	}
	if start < len(a) {
		result = append(result, a[start:])
	}
	return result
}

func contains(a []string, s string) bool {
	for _, c := range s {
		if slices.Contains(a, string(c)) {
			return true
		}
	}
	return false
}

func separateParenthesis(a string) []string {
	var result []string
	start := 0
	for i, c := range a {
		if c == '(' || c == ')' || c == '[' || c == ']' || c == '{' || c == '}' {
			if i > start {
				result = append(result, a[start:i])
			}
			result = append(result, string(c))
			start = i + 1
		}
	}
	if start < len(a) {
		result = append(result, a[start:])
	}
	return result
}

func Label(split_code [][]string) [][]string {
	//if the keyword that is found is anx, label as keyword and the next as identifier
	var new_words [][]string
	keyWords := []string{"IDENTIFIER", ".>", "anx", "if", "elif", "else"}
	parenthesis := []string{"(", ")", "[", "]", "{", "}"}
	quotes := []string{`'`, `"`}
	for _, w := range split_code {
		//other checks before you conclude with the identifier section

		if !slices.Contains(keyWords, w[0]) && !slices.Contains(parenthesis, w[0]) {
			w = append([]string{"IDENTIFIER"}, w...)
		}
		if slices.Contains(parenthesis, w[0]) {
			switch w[0] {
			case "(":
				w = append([]string{"OPENNING PARENTHESIS"}, w...)
			case ")":
				w = append([]string{"CLOSING PARENTHESIS"}, w...)
			case "[":
				w = append([]string{"OPENNING SQUARE BRACKET"}, w...)
			case "]":
				w = append([]string{"CLOSING SQUARE BRACKET"}, w...)
			case "{":
				w = append([]string{"OPENNING CURLY BRACE"}, w...)
			case "}":
				w = append([]string{"CLOSING CURLY BRACE"}, w...)
			}
		}
		if slices.Contains(quotes, w[0]) {
			switch w[0] {
			case `'`:
				w = append([]string{"SINGLE QUOTE"}, w...)
			case `"`:
				w = append([]string{"DOUBLE QUOTE"}, w...)
			}
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
		words = Sub_split(words)
		for _, w := range words {
			l.Split_code = append(l.Split_code, []string{w})
		}
	}
	l.Split_code = Label(l.Split_code)
}
