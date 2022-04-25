package piscine

import "ft"

func PrintDigits() {
	//for文で書く
	for c:= '0'; c <= '9';c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}