package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("brainfuck/test-h.b")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	buff := make([]byte, 256)
	for {
		c, _ := input.Read(buff)
		if c == 0 {
			break
		}
		fmt.Print(string(buff[:c]))
	}
	input.Close()
}
