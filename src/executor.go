package gostack

import "strconv"

type Executor struct {
	stack    *Stack
	elements []ASTElement
}

func NewExecutor(elements []ASTElement) *Executor {
	p := new(Executor)
	p.stack = NewStack()
	p.elements = elements
	return p
}

func (e *Executor) RunIntoString() string {
	result := ""
	for _, element := range e.elements {
		expr, isExpr := element.(ASTExpr)
		if isExpr {
			switch expr.ExprKind {
			case EXPR_KIND_ADD:
				op1 := e.stack.Pop()
				op2 := e.stack.Pop()
				e.stack.Push(op1 + op2)
			case EXPR_KIND_SUB:
				op1 := e.stack.Pop()
				op2 := e.stack.Pop()
				e.stack.Push(op2 - op1)
			case EXPR_KIND_DROP:
				e.stack.Pop()
			case EXPR_KIND_DUP:
				e.stack.Push(e.stack.Top())
			case EXPR_KIND_PRINT:
				result += " " + strconv.Itoa(int(e.stack.Top()))
			default:
			}
		} else {
			unaryExpr, isUnaryExpr := element.(ASTUnaryExpr)
			if isUnaryExpr {
				if unaryExpr.UnaryExprKind == UNARY_EXPR_KIND_PUSH {
					e.stack.Push(unaryExpr.Operand)
				}
			}
		}
	}

	return result
}
