package gostack

import "testing"

func TestAstElement(t *testing.T) {
	addExpr := ASTExpr{ExprKind: EXPR_KIND_ADD}
	if addExpr.ExprKind != EXPR_KIND_ADD {
		t.Errorf("addExpr.ExprKind == %d, want %d", addExpr.ExprKind, EXPR_KIND_ADD)
	}
	pushExpr := NewPushExpr(1)
	if pushExpr.ExprKind != EXPR_KIND_UNARY {
		t.Errorf("pushExpr.ExprKind == %d, want %d", pushExpr.ExprKind, EXPR_KIND_UNARY)
	}
	if pushExpr.UnaryExprKind != UNARY_EXPR_KIND_PUSH {
		t.Errorf("pushExpr.UnaryExprKind == %d, want %d", pushExpr.ExprKind, UNARY_EXPR_KIND_PUSH)
	}
	if pushExpr.Operand != 1 {
		t.Errorf("pushExpr.Operand == %d, want %d", pushExpr.Operand, 1)
	}
}
