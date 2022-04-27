package piscine

func isDigit(c rune) bool {
	if '0' <= c && c <= '9'{
		return true
	}
	return false
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(isDigit(c)) {
			return false
		}
	}
	return true
}