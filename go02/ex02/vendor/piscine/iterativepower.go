package piscine


func IterativePower(nb int, power int) int {
	var ret int = 1;


	if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	}
	for ;power > 0; power-- {
		ret *= nb
	}
	return ret
}