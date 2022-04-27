package piscine

import (
	"ft"
	"os"
)

func printprogramname() {
	cmdline := os.Args[0]
	for i,c := range cmdline {
		if i > 1 {
			ft.PrintRune(c)
		}
	}
	ft.PrintRune('\n')
}