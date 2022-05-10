package piscine

func FirstRune(s string) rune {
	if s == "" {
		return 0
	}
	s_slice := []rune(s)

	return s_slice[0]
}