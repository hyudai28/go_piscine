package piscine


func IterativeFactorial(nb int) int {
	var factorial int = 1

	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb >= 21 {
		return 0
	}
	for ;nb > 0; nb-- {
		factorial *= nb
	}
	return factorial
}