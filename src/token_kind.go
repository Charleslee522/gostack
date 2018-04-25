package gostack

type TokenKind uint16

const (
	PUSH TokenKind = 0 + iota
	ADD
	SUB
	DROP
	DUP
	PRINT
	NUM
	EOF
)

func (kind TokenKind) String() string {
	names := [...]string{
		"PUSH",
		"ADD",
		"SUB",
		"DROP",
		"DUP",
		"PRINT",
		"NUM",
		"EOF"}
	if kind < PUSH || kind > EOF {
		return "UNKNOWN"
	}
	return names[kind]
}
