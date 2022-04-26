package piscine

func StrLen(s string) int {
	var str_length int = 0;
	for i, _ := range s {
		str_length = i
	}
	return (str_length)
}

func StrRev(s string) string {
	var s_i int
	rev  := []rune(s)
	tmp  := []rune(s)

	s_i = StrLen(s)
	for i:= 0; i <= s_i; i++ {
		rev[i] = tmp[s_i - i]
	}
	return ((string)(rev))
}