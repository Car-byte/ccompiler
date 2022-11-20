package ccompiler

type Token struct {
	TokenType TokenType
	Next      *Token
	Val       int
	Location  *byte
	Len       int
}

func (token Token) GetVal() int {
	if token.TokenType != NUM_LIT {
		panic("Expected a number")
	}

	return token.Val
}

func CreateToken(tokenType TokenType, start *byte, len int) Token {
	token := Token{}
	token.TokenType = tokenType
	token.Location = start
	token.Len = len
	return token
}

func EqualToken(firstToken, secondToken Token) bool {
	if firstToken.TokenType != secondToken.TokenType {
		return false
	}

	if secondToken.Len != secondToken.Len {
		return false
	}

	if *firstToken.Location != *secondToken.Location {
		return false
	}

	return true
}
