package gostack

import (
	"testing"
)

func TestStackBasic(t *testing.T) {
	stack := NewStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	result := stack.Top()
	Check(t, result, 3)
	result = stack.Pop()
	Check(t, result, 3)
	result = stack.Pop()
	Check(t, result, 2)
	result = stack.Pop()
	Check(t, result, 1)
}
