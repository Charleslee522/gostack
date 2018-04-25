package gostack

import (
	"strconv"
	. "strings"
)

type Tokenizer struct {
	tokens []TokenKind
}

func NewTokenizer(data string) *Tokenizer {
	p := Tokenizer{}
	stringTokens := Split(data, " ")

	for _, data := range stringTokens {
		switch data {
		case "PUSH":
			p.tokens = append(p.tokens, PUSH)
		case "ADD":
			p.tokens = append(p.tokens, ADD)
		case "SUB":
			p.tokens = append(p.tokens, SUB)
		case "DROP":
			p.tokens = append(p.tokens, DROP)
		case "DUP":
			p.tokens = append(p.tokens, DUP)
		case "PRINT":
			p.tokens = append(p.tokens, PRINT)
		default:
			if _, err := strconv.Atoi(data); err == nil {
				p.tokens = append(p.tokens, NUM)
			}

		}
	}

	return &p
}

func (t *Tokenizer) GetToken() TokenKind {
	token := t.tokens[0]
	t.tokens = t.tokens[1:]
	return token
}
