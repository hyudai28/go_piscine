package piscine

func RecursiveFunc(nb int) int {
	if (nb == 1) {
		return (1)
	}
	nb *= RecursiveFunc(nb - 1)
	return nb
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb >= 21 {
		return 0
	}
	return RecursiveFunc(nb)
}
