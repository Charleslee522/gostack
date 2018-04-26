package gostack

// func TestExecutor1(t *testing.T) {
// 	tokenizer := NewTokenizer("PUSH 1 PUSH 2 ADD PRINT")
// 	stack := NewStack(tokenizer)
// 	executor := Executor(stack)
// 	result := executor.RunIntoString()
// 	if result != "3" {
// 		t.Errorf("result == %s, want %s", result, "3")
// 	}
// }

// func TestExecutor2(t *testing.T) {
// 	tokenizer := NewTokenizer("PUSH 1 PUSH 2 SUB PRINT")
// 	stack := NewStack(tokenizer)
// 	executor := Executor(stack)
// 	executor.Run()
// 	result := executor.RunIntoString()
// 	if result != "1" {
// 		t.Errorf("result == %s, want %s", result, "3")
// 	}
// }

// func TestExecutor3(t *testing.T) {
// 	tokenizer := NewTokenizer("PUSH 1 PUSH 2 DROP PRINT DUP PRINT ADD PRINT")
// 	stack := NewStack(tokenizer)
// 	executor := Executor(stack)
// 	executor.Run()
// 	result := executor.RunIntoString()
// 	if result != "1 1 2" {
// 		t.Errorf("result == %s, want %s", result, "1 1 2")
// 	}
// }
