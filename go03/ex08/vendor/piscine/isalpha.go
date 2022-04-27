package piscine

func isDigit(c rune) bool {
	if '0' <= c && c <= '9'{
		return true
	}
	return false
}

func isAlpha(c rune) bool {
	if (('A' <= c && c <= 'Z')  || ('a' <= c && c <= 'z')){
		return true
	}
	return false
}

func IsAlpha(s string) bool {
	for _, c := range s {
		if !(isAlpha(c) || isDigit(c)) {
			return false
		}
	}
	return true
}