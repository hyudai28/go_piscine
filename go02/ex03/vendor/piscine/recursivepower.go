package piscine

func RecursivePower(nb int, power int) int {
	var ret int = 1;


	if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	}
	//for ;power > 0; power-- {
	//	ret *= nb
	//}
	ret *= RecursivePower(nb, power - 1) * nb
	return ret
}