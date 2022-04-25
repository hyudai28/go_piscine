package piscine

import "ft"

func PrintReverseAlphabet() {
	//for文で書く
	for c:= 'z'; c >= 'a';c-- {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}