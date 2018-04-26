package gostack

import "testing"

func TestTokenizer(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 ADD PRINT")
	tok := tokenizer.GetToken()
	if tok.Kind != PUSH {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != NUM {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, NUM.String(), NUM)
	}
	if tok.Value != 1 {
		t.Errorf("tok.Value == %d, want %d", tok.Value, 1)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PUSH {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != NUM {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, NUM.String(), NUM)
	}
	if tok.Value != 2 {
		t.Errorf("tok.Value == %d, want %d", tok.Value, 2)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != ADD {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, ADD.String(), ADD)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PRINT {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PRINT.String(), PRINT)
	}
}

func TestTokenizer2(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 SUB PRINT")
	tok := tokenizer.GetToken()
	if tok.Kind != PUSH {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != NUM {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PUSH {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != NUM {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != SUB {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, SUB.String(), SUB)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PRINT {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PRINT.String(), PRINT)
	}
}

func TestTokenizer3(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 DROP PRINT DUP PRINT ADD PRINT")
	tok := tokenizer.GetToken()
	if tok.Kind != PUSH {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != NUM {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PUSH {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != NUM {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != DROP {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, DROP.String(), DROP)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PRINT {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PRINT.String(), PRINT)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != DUP {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, DUP.String(), DUP)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PRINT {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PRINT.String(), PRINT)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != ADD {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, ADD.String(), ADD)
	}
	tok = tokenizer.GetToken()
	if tok.Kind != PRINT {
		t.Errorf("tok.Kind == %d, want %s %d", tok.Kind, PRINT.String(), PRINT)
	}
}
