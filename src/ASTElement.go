package gostack

type ASTExprKind uint16

const (
	EXPR_KIND_NONE ASTExprKind = 0 + iota
	EXPR_KIND_ADD
	EXPR_KIND_SUB
	EXPR_KIND_DROP
	EXPR_KIND_DUP
	EXPR_KIND_PRINT
	EXPR_KIND_UNARY
)

type ASTUnaryExprKind uint16

const (
	UNARY_EXPR_KIND_PUSH ASTUnaryExprKind = 0 + iota
	UNARY_EXPR_KIND_NONE
)

type ASTElement interface {
}

type ASTExpr struct {
	ASTElement
	ExprKind ASTExprKind
}

type ASTUnaryExpr struct {
	ASTExpr
	UnaryExprKind ASTUnaryExprKind
	Operand       ElementType
}

func NewExpr(tok Token) *ASTExpr {
	p := new(ASTExpr)
	p.ASTElement = new(ASTElement)
	switch tok.Kind {
	case ADD:
		p.ExprKind = EXPR_KIND_ADD
	case SUB:
		p.ExprKind = EXPR_KIND_SUB
	case DROP:
		p.ExprKind = EXPR_KIND_DROP
	case DUP:
		p.ExprKind = EXPR_KIND_DUP
	case PRINT:
		p.ExprKind = EXPR_KIND_PRINT
	default:
		p.ExprKind = EXPR_KIND_NONE
	}
	return p
}

func NewPushExpr(value ElementType) *ASTUnaryExpr {
	p := new(ASTUnaryExpr)
	p.ASTExpr = ASTExpr{ASTElement: new(ASTElement), ExprKind: EXPR_KIND_UNARY}
	p.UnaryExprKind = UNARY_EXPR_KIND_PUSH
	p.Operand = value
	return p
}
