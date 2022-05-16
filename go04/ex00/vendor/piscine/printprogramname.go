package piscine

import (
	"ft"
	"os"
)

func Strlen(s string) int {
	var str_length int = 0
	for range s {
		str_length++
	}
	return str_length
}

func Printprogramname() {
	cmdline := os.Args[0]
	name_slice := []rune(cmdline)

	name_length := Strlen(cmdline) - 1
	if name_length == -1 {
		return 
	}
	for name_length >= 0 {
		if name_slice[name_length] == '/' {
			name_length++
			break
		}
		name_length--
	}
	if name_length == -1 {
		name_length++
	}
	for _, r := range name_slice[name_length:] {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}