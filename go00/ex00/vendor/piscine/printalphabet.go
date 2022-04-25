package piscine

import "ft"

func PrintAlphabet() {
	//for文で書く
	for c:= 'a'; c <= 'z';c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}