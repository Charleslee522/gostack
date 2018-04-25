package gostack

import "testing"

func TestTokenizer(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 ADD PRINT")
	tok := tokenizer.GetToken()
	if tok != PUSH {
		t.Errorf("tok == %d, want %s %d", tok, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok != NUM {
		t.Errorf("tok == %d, want %s %d", tok, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok != PUSH {
		t.Errorf("tok == %d, want %s %d", tok, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok != NUM {
		t.Errorf("tok == %d, want %s %d", tok, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok != ADD {
		t.Errorf("tok == %d, want %s %d", tok, ADD.String(), ADD)
	}
	tok = tokenizer.GetToken()
	if tok != PRINT {
		t.Errorf("tok == %d, want %s %d", tok, PRINT.String(), PRINT)
	}
}

func TestTokenizer2(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 SUB PRINT")
	tok := tokenizer.GetToken()
	if tok != PUSH {
		t.Errorf("tok == %d, want %s %d", tok, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok != NUM {
		t.Errorf("tok == %d, want %s %d", tok, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok != PUSH {
		t.Errorf("tok == %d, want %s %d", tok, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok != NUM {
		t.Errorf("tok == %d, want %s %d", tok, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok != SUB {
		t.Errorf("tok == %d, want %s %d", tok, SUB.String(), SUB)
	}
	tok = tokenizer.GetToken()
	if tok != PRINT {
		t.Errorf("tok == %d, want %s %d", tok, PRINT.String(), PRINT)
	}
}

func TestTokenizer3(t *testing.T) {
	tokenizer := NewTokenizer("PUSH 1 PUSH 2 DROP PRINT DUP PRINT ADD PRINT")
	tok := tokenizer.GetToken()
	if tok != PUSH {
		t.Errorf("tok == %d, want %s %d", tok, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok != NUM {
		t.Errorf("tok == %d, want %s %d", tok, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok != PUSH {
		t.Errorf("tok == %d, want %s %d", tok, PUSH.String(), PUSH)
	}
	tok = tokenizer.GetToken()
	if tok != NUM {
		t.Errorf("tok == %d, want %s %d", tok, NUM.String(), NUM)
	}
	tok = tokenizer.GetToken()
	if tok != DROP {
		t.Errorf("tok == %d, want %s %d", tok, DROP.String(), DROP)
	}
	tok = tokenizer.GetToken()
	if tok != PRINT {
		t.Errorf("tok == %d, want %s %d", tok, PRINT.String(), PRINT)
	}
	tok = tokenizer.GetToken()
	if tok != DUP {
		t.Errorf("tok == %d, want %s %d", tok, DUP.String(), DUP)
	}
	tok = tokenizer.GetToken()
	if tok != PRINT {
		t.Errorf("tok == %d, want %s %d", tok, PRINT.String(), PRINT)
	}
	tok = tokenizer.GetToken()
	if tok != ADD {
		t.Errorf("tok == %d, want %s %d", tok, ADD.String(), ADD)
	}
	tok = tokenizer.GetToken()
	if tok != PRINT {
		t.Errorf("tok == %d, want %s %d", tok, PRINT.String(), PRINT)
	}
}
