package main

import (
	"fmt"
	"os"
	"strconv"
)

func nextInt(input string, i int) (int, int) {
	start := i

	for i < len(input) && input[i] <= '9' && input[i] >= '0' {
		i++
	}

	val, err := strconv.Atoi(input[start:i])
	if err != nil {
		fmt.Printf("Unable to parse: %s\n", input[start:i])
		os.Exit(1)
	}

	return val, i
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("%s: invalid number of args", args[0])
		os.Exit(1)
	}

	input := args[1]

	fmt.Println("  .globl main")
	fmt.Println("main:")

	val, i := nextInt(input, 0)

	fmt.Printf("  mov $%d, %%rax\n", val)

	for i < len(input) {
		if input[i] == '+' {
			val, i = nextInt(input, i+1)
			fmt.Printf("  add $%d, %%rax\n", val)
		} else if input[i] == '-' {
			val, i = nextInt(input, i+1)
			fmt.Printf("  sub $%d, %%rax\n", val)
		} else {
			fmt.Printf("%s: unable to parse %s", args[0], args[1])
			os.Exit(1)
		}
	}

	fmt.Println("  ret")
}
