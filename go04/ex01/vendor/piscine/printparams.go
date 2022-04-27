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

func printparams() {
    for i, v := range os.Args {
		if i >= 1 {
			printLine(v)
		}
    }
}