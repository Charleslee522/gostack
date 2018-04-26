package gostack

import (
	"strconv"
	"strings"
	. "strings"
)

type Token struct {
	Kind  TokenKind
	Value ElementType
}

type Tokenizer struct {
	tokens []Token
}

func NewTokenizer(data string) *Tokenizer {
	p := new(Tokenizer)
	s := strings.TrimSpace(data)
	stringTokens := Split(s, " ")

	for _, data := range stringTokens {
		switch data {
		case "PUSH":
			p.tokens = append(p.tokens, Token{PUSH, 0})
		case "ADD":
			p.tokens = append(p.tokens, Token{ADD, 0})
		case "SUB":
			p.tokens = append(p.tokens, Token{SUB, 0})
		case "DROP":
			p.tokens = append(p.tokens, Token{DROP, 0})
		case "DUP":
			p.tokens = append(p.tokens, Token{DUP, 0})
		case "PRINT":
			p.tokens = append(p.tokens, Token{PRINT, 0})
		default:
			if value, err := strconv.Atoi(data); err == nil {
				p.tokens = append(p.tokens, Token{NUM, ElementType(value)})
			}
		}
	}
	p.tokens = append(p.tokens, Token{EOF, 0})

	return p
}

func (t *Tokenizer) GetToken() Token {
	token := t.tokens[0]
	t.tokens = t.tokens[1:]
	return token
}
