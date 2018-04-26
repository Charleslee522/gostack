package gostack

import "testing"

func TestExecutor1(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 ADD PRINT")
	parser := Parser{tokenizer}
	elements := parser.Parse()
	executor := NewExecutor(elements)
	result := executor.RunIntoString()
	if result != " 3" {
		t.Errorf("result == %s, want %s", result, " 3")
	}
}

func TestExecutor2(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 SUB PRINT")
	parser := Parser{tokenizer}
	elements := parser.Parse()
	executor := NewExecutor(elements)
	result := executor.RunIntoString()
	if result != " -1" {
		t.Errorf("result == %s, want %s", result, " -1")
	}
}

func TestExecutor3(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 DROP PRINT DUP PRINT ADD PRINT")
	parser := Parser{tokenizer}
	elements := parser.Parse()
	executor := NewExecutor(elements)
	result := executor.RunIntoString()
	if result != " 1 1 2" {
		t.Errorf("result == %s, want %s", result, " 1 1 2")
	}
}

func TestExecutor4(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 PUSH 3 PUSH 4 PUSH 5 DROP PRINT DUP ADD SUB SUB PRINT")
	parser := Parser{tokenizer}
	elements := parser.Parse()
	executor := NewExecutor(elements)
	result := executor.RunIntoString()
	if result != " 4 7" {
		t.Errorf("result == %s, want %s", result, " 4 7")
	}
}
