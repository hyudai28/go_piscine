package piscine

import "ft"

func PrintComb() {
	var	num_1 = 2
	var num_2 = 1
	var num_3 = 0
	for num_3 <= 7 {
		for num_2 <= 8 {
			for num_1 <= 9 {
				ft.PrintRune((rune)(num_3 + '0'))
				ft.PrintRune((rune)(num_2 + '0'))
				ft.PrintRune((rune)(num_1 + '0'))
				if !(num_3 == 7 && num_2 == 8 && num_1 == 9) {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
				num_1++
			}
			num_2++
			num_1 = num_2 + 1
		}
		num_3++
		num_2 = num_3 + 1
		num_1 = num_2 + 1
	}
	ft.PrintRune('\n')
}