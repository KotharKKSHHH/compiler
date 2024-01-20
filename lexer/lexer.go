package lexer

import (
	"fmt"
	"regexp"
)

type TokenType int

const (
	Identifier TokenType = iota
	Number
	Operator
	Whitespace
	Keyword
	Punctuation
)

type Token struct {
	Type  TokenType
	Value string
}

func Lex(input string) []Token {
	var tokens []Token

	tokenPatterns := []struct {
		Type    TokenType
		Pattern string
	}{
		{Identifier, `[a-zA-Z_][a-zA-Z0-9_]*`},
		{Number, `[0-9]+`},
		{Operator, `[+\-*/]`},
		{Whitespace, `\s+`},
		{Keyword, `def|int|vector|var|for|to|length|return`},
		{Punctuation, `[{}();=]`},
	}

	for len(input) > 0 {
		match := false

		for _, pattern := range tokenPatterns {
			re := regexp.MustCompile("^" + pattern.Pattern)
			if submatch := re.FindString(input); submatch != "" {
				tokens = append(tokens, Token{Type: pattern.Type, Value: submatch})
				input = input[len(submatch):]
				match = true
				break
			}
		}

		if !match {
			fmt.Printf("Invalid token: %v\n", input[0])
			input = input[1:]
		}
	}

	return tokens
}
