package piscine

import "ft"

func count_right(left int) {
	for right := left + 1; right <= 99 ; right++ {
		ft.PrintRune((rune)(left / 10 + '0'))
		ft.PrintRune((rune)(left % 10 + '0'))
		ft.PrintRune(' ')
		ft.PrintRune((rune)(right / 10 + '0'))
		ft.PrintRune((rune)(right % 10 + '0'))
		if !(left == 98 && right == 99) {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
	}
}

func PrintComb2() {
	for left := 0; left <= 98; left++ {
		count_right(left)
	}
}