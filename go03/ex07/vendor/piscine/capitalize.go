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

func Capitalize(s string) string {
	lower_s := []rune(s)
	var before_c bool = true

	for i, c := range lower_s {
		if before_c == true {
			if 'a' <= c && c <= 'z' {
				lower_s[i] -=  32
			}
		}
		if (isAlpha(c) || isDigit(c)) {
			before_c = false
		} else {
			before_c = true
		}
	}
	return (string)(lower_s)
}