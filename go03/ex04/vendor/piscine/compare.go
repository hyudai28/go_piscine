package piscine

func Strlen(s string) int {
	var str_length int = 0
	for i, _ := range s {
		str_length = i
	}
	return str_length + 1
}

func Compare(a, b string) int {
	var a_length = Strlen(a)
	var b_length = Strlen(b)
	a_slice := a[:]
	b_slice := b[:]

	for i := 0; i < a_length; i++ {
		if i == b_length && a_length != b_length {
			return 1
		}
		if a_slice[i] > b_slice[i] {
			return 1
		} else if a_slice[i] < b_slice[i]{
			return -1
		}

	}
	return 0
}