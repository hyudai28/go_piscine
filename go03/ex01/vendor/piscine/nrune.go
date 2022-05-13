package piscine

//import "fmt"
func Strlen(s string) int {
	var str_length int = 0
	for range s {
		str_length++
	}
	return str_length
}

func NRune(s string, n int) rune {
	s_slice := []rune(s)

	var str_length = Strlen(s)
	if (str_length == 0) {
		return 0
	}
	if (n > 0 && n <= str_length) {
		return s_slice[n - 1]
	}
	return 0
}