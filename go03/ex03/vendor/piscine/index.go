package piscine

func Strlen(s string) int {
	var str_length int = 0
	for range s {
		str_length++
	}
	return str_length
}

func Index(s string, toFind string) int {
	var s_length = Strlen(s)
	var toFind_length = Strlen(toFind)

	for i := 0; i + toFind_length <= s_length; i++ {
		s_slice := s[i: i + toFind_length]
		if s_slice == toFind {
			return i
		}
	}
	return -1
}