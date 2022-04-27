package piscine

//import "fmt"
func Strlen(s string) int {
	var str_length int = 0
	for i, _ := range s {
		str_length = i
	}
	return str_length
}

func NRune(s string, n int) rune {
	s_slice := []rune(s)

	var str_length = Strlen(s)
	n -= 1
	if (n > 0 && n <= str_length) {
		return s_slice[n]
	}
	return 0
}