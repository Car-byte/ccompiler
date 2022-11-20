package main

import (
	. "ccompiler/ccompiler"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func nextInt(input []byte, i int) (int, int) {
	start := i

	for i < len(input) && isDigit(input[i]) {
		i++
	}

	val, err := strconv.Atoi(string(input[start:i]))
	if err != nil {
		fmt.Printf("Unable to parse: %s\n", input[start:i])
		os.Exit(1)
	}

	return val, i - start
}

func isDigit(b byte) bool {
	return b <= byte('9') && b >= byte('0')
}

func tokenize(input []byte) *Token {
	head := Token{}
	cur := &head

	for i := 0; i < len(input); {
		if unicode.IsSpace(rune(input[i])) {
			i++
		} else if isDigit(input[i]) {
			val, len := nextInt(input, i)

			nextToken := CreateToken(NUM_LIT, &input[i], len)
			cur.Next = &nextToken
			cur.Next.Val = val
			cur = cur.Next

			i += len
		} else if input[i] == byte('+') || input[i] == byte('-') {
			nextToken := CreateToken(PUNCTUATOR, &input[i], 1)
			cur.Next = &nextToken
			cur = cur.Next

			i++
		} else {
			panic("Error unable to create token")
		}
	}

	eofToken := CreateToken(EOF, &input[len(input)-1], 0)
	cur.Next = &eofToken

	return head.Next
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("%s: invalid number of args", args[0])
		os.Exit(1)
	}

	curToken := tokenize([]byte(args[1]))

	fmt.Println("  .globl main")
	fmt.Println("main:")

	fmt.Printf("  mov $%d, %%rax\n", curToken.GetVal())
	curToken = curToken.Next

	minusSign := byte('-')
	plusSign := byte('+')
	subtractToken := CreateToken(PUNCTUATOR, &minusSign, 1)
	additionToken := CreateToken(PUNCTUATOR, &plusSign, 1)

	for curToken.TokenType != EOF {
		if EqualToken(subtractToken, *curToken) {
			fmt.Printf("  sub $%d, %%rax\n", curToken.Next.GetVal())
		} else if EqualToken(additionToken, *curToken) {
			fmt.Printf("  add $%d, %%rax\n", curToken.Next.GetVal())
		} else {
			fmt.Printf("%s: unable to parse %s", args[0], args[1])
			os.Exit(1)
		}

		curToken = curToken.Next.Next
	}

	fmt.Println("  ret")
}
