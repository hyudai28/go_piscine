package piscine

import (
	"ft"
	"os"
)

func printLine(line string) {
	for _,c := range line {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func PrintArray(args []string) {
	args_len := 0
	for i, _ := range args {
		args_len = i
	}
	for ; args_len > 0; args_len-- {
		printLine(args[args_len])
	}
}

func RevParams() {
	PrintArray(os.Args)
}