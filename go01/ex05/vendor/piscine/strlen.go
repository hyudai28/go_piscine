package piscine


func StrLen(s string) int {
	var str_length int = 0;
	for i, _ := range s {
		str_length = i
	}
	return (str_length + 1)
}