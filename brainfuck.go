package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("brainfuck/hoge.b")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pc := 0
	mem := make([]byte, 30000)
	buff := make([]byte, 30000)
	for {
		c, _ := input.Read(buff)
		if c == 0 {
			break
		}
		// str := string(buff[:c])
		// // fmt.Print(str)
		// if str == "-" {
		// 	fmt.Print(str)
		// }
		switch str {
		case "+":
			mem[pc] += 1
		case "-":
			mem[pc] -= 1
		case "<":
			pc += 1
		case ">":
			pc -= 1
		case "[":
		case "]":
		case ",":
		case ".":
			fmt.Printf("%s", mem[pc])
		default:
		}
	}
	input.Close()
}
