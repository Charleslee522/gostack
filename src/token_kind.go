package gostack

type TokenKind uint16

const (
	NONE TokenKind = 0 + iota
	PUSH
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
		"NONE",
		"PUSH",
		"ADD",
		"SUB",
		"DROP",
		"DUP",
		"PRINT",
		"NUM",
		"EOF"}
	if kind < NONE || kind > EOF {
		return "UNKNOWN"
	}
	return names[kind]
}
