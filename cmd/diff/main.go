package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Chara-X/diff"
)

func main() {
	var from, _ = os.ReadFile(os.Args[1])
	var to, _ = os.ReadFile(os.Args[2])
	var changes = diff.Diff(strings.Split(string(from), "\n"), strings.Split(string(to), "\n"))
	for _, v := range changes {
		switch v.Kind {
		case 0:
			fmt.Printf("\x1b[32m%v: %v\n\x1b[0m", v.Position, v.To)
		case 1:
			fmt.Printf("\x1b[31m%v: %v\n\x1b[0m", v.Position, v.From)
		case 2:
			fmt.Printf("\x1b[33m%v: %v -> %v\n\x1b[0m", v.Position, v.From, v.To)
		}
	}
}
