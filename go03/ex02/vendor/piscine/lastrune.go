package piscine


func Strlen(s string) int {
	var str_length int = 0
	for i, _ := range s {
		str_length = i
	}
	return str_length
}

func LastRune(s string) rune {
	s_slice := []rune(s)

	var str_length = Strlen(s)
	return s_slice[str_length]
}