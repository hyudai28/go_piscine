package piscine

//func StrLen(s string) int {
//	var str_length int = 0;
//	for i, _ := range s {
//		str_length = i
//	}
//	return (str_length)
//}

//func StrRev(s string) string {
//	var s_i int
//	rev  := []rune(s)
//	tmp  := []rune(s)

//	s_i = StrLen(s)
//	if (s_i == 0) {
//		return s
//	}
//	for i:= 0; i <= s_i; i++ {
//		rev[i] = tmp[s_i - i]
//	}
//	return ((string)(rev))
//}


package piscine

func StrRev(s string) string {
    var r_str string
    for _, c := range s {
        top := string(c)
        top += r_str
        r_str = top
    }
    return r_str
}