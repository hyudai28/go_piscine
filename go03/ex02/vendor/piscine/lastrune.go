package piscine


func Strlen(s string) int {
	var str_length int = 0
	for range s {
		str_length++
	}
	return str_length
}

func LastRune(s string) rune {
	s_slice := []rune(s)

	var str_length = Strlen(s)
	if str_length == 0 {
		return 0
	}
	str_length -= 1
	return s_slice[str_length]
}