package gostack

type ASTExprKind uint16

const (
	EXPR_KIND_ADD ASTExprKind = 0 + iota
	EXPR_KIND_SUB
	EXPR_KIND_DROP
	EXPR_KIND_DUP
	EXPR_KIND_PRINT
	EXPR_KIND_UNARY
	NONE
)

type ASTUnaryExprKind uint16

const (
	UNARY_EXPR_KIND_PUSH ASTUnaryExprKind = 0 + iota
	UNARY_EXPR_KIND_NONE
)

type ASTElement struct {
}

type ASTExpr struct {
	ASTElement
	ExprKind ASTExprKind
}

type ASTUnaryExpr struct {
	ASTExpr
	UnaryExprKind ASTUnaryExprKind
	Operand       int
}

func NewPushExpr(value int) *ASTUnaryExpr {
	p := new(ASTUnaryExpr)
	p.ASTExpr = ASTExpr{ExprKind: EXPR_KIND_UNARY}
	p.UnaryExprKind = UNARY_EXPR_KIND_PUSH
	p.Operand = value
	return p
}
