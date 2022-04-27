package piscine

func Sqrt(nb int) int {
	for i := 1; i * 2 <= nb;i++ {
		if i * i == nb {
			return i
		}
	}
	return 0
}
