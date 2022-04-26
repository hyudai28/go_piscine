package piscine

func Atoi(s string) int {
	var ret int = 0
	var plus_minus int = 1

	for i, c := range s {
		if (i == 0) {
			if (c == '+') {
				plus_minus = 1
			} else if (c == '-') {
				plus_minus = -1
			} else if !(c >= '0' && c <= '9') {
				return (0)
			} else {
				ret +=  (int)(c - '0')
			}
		} else if !(c >= '0' && c <= '9') {
			return (0)
		} else {
		ret *= 10
		ret += (int)(c - '0')
		}
	}
	return (ret * plus_minus)
}