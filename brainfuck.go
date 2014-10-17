package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("brainfuck/33.b")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	reg := 0
	mem := make([]int, 30000)
	buff := make([]byte, 30000)
	for {
		c, _ := input.Read(buff)
		if c == 0 {
			break
		}
		str := string(buff[:c])
		for pc := 0; pc < c; pc++ {
			switch string(str[pc]) {
			case "+":
				mem[reg] += 1
			case "-":
				mem[reg] -= 1
			case "<":
				reg += 1
			case ">":
				reg -= 1
			case "[":
				// if mem[reg] != 0 {
				// 	break
				// }
				// nest := 0
				// for pc < str.length {
				// 	if string(str[pc]) == "[" {
				// 		nest += 1
				// 	} else if string(str[pc]) == "]" {
				// 		nest -= 1
				// 		if nest == 0 {
				// 			break
				// 		}
				// 	}
				// 	pc += 1
				// }
			case "]":
				// if mem[reg] == 0 {
				// 	break
				// }
				// nest := 0
				// for {
				// 	if reg <= 0 {
				// 		break
				// 	}
				// 	if string(str[reg]) == "]" {
				// 		nest += 1
				// 	} else if string(str[reg]) == "[" {
				// 		nest -= 1
				// 		if nest == 0 {
				// 			break
				// 		}
				// 	}
				// 	reg -= 1
				// }
			case ",":
			case ".":
				fmt.Println(string(mem[reg]))
			default:
			}
		}
	}
	input.Close()
}
