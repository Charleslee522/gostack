package gostack

import (
	"testing"
)

func TestParser(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 ADD PRINT")
	parser := Parser{tokenizer}
	elements := parser.Parse()

	idx := 0
	unaryExpr, isUnaryExpr := elements[idx].(ASTUnaryExpr)
	if !isUnaryExpr {
		t.Errorf("%d'th element is not UnaryExpr, want UnaryExpr", idx)
	} else {
		correct := NewPushExpr(1)
		if unaryExpr.Operand != correct.Operand {
			t.Errorf("%d'th element == %q, want %q", idx, unaryExpr, correct)
		}
	}

	idx = 2
	expr, isExpr := elements[idx].(ASTExpr)
	if !isExpr {
		t.Errorf("%d'th element is not expr, want expr", idx)
	} else {
		if expr.ExprKind != EXPR_KIND_ADD {
			t.Errorf("%d'th element == %d, want %d", idx, expr.ExprKind, EXPR_KIND_ADD)
		}
	}

	idx = 3
	expr, isExpr = elements[idx].(ASTExpr)
	if !isExpr {
		t.Errorf("%d'th element is not expr, want expr", idx)
	} else {
		if expr.ExprKind != EXPR_KIND_PRINT {
			t.Errorf("%d'th element == %d, want %d", idx, expr.ExprKind, EXPR_KIND_PRINT)
		}
	}
}
