package gostack

type Parser struct {
	tokenizer *Tokenizer
}

func NewParser(tokenizer *Tokenizer) *Parser {
	p := new(Parser)
	p.tokenizer = tokenizer
	return p
}

func (p *Parser) Parse() []ASTElement {
	result := make([]ASTElement, 0)
	for {
		tok := p.tokenizer.GetToken()
		switch tok.Kind {
		case PUSH:
			tok = p.tokenizer.GetToken()
			if tok.Kind != NUM {
				//exception
			} else {
				result = append(result, *NewPushExpr(tok.Value))
			}
		case ADD:
			result = append(result, *NewExpr(tok))
		case SUB:
			result = append(result, *NewExpr(tok))
		case DROP:
			result = append(result, *NewExpr(tok))
		case DUP:
			result = append(result, *NewExpr(tok))
		case PRINT:
			result = append(result, *NewExpr(tok))
		default:
		}
		if tok.Kind == EOF {
			break
		}
	}
	return result
}
