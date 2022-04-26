package piscine

func BasicAtoi2(s string) int {
	var ret int = 0

	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			return (0)
		}
		ret *= 10
		ret += (int)(c - '0')
	}
	return (ret)
}