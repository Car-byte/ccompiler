package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("%s: invalid number of args", args[0])
		os.Exit(1)
	}

	fmt.Println("  .globl main")
	fmt.Println("main:")

	argVal, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%s: unable to parse input %s", args[0], args[1])
		os.Exit(1)
	}

	fmt.Printf("  mov $%d, %%rax\n", argVal)
	fmt.Println("  ret")
}
