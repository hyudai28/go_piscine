package piscine

import "ft"

func printLine(line string) {
	for _,c := range line {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	for _, s := range a {
			printLine(s)
	}
}