package ccompiler

type TokenType int

const (
	PUNCTUATOR TokenType = iota
	NUM_LIT
	EOF
)
