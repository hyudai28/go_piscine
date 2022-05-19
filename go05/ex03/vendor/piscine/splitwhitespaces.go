package piscine

//import "fmt"

func isSpace(c rune) int {
	if c == ' ' || c == '\n' || c == '\t' || c == '\v' {
		return 1;
	}
	return 0
}

func Split_count(s string) int {
	var array_count int
	space_switch := true

	for _,c := range s {
		//if c == ' ' {
		if isSpace(c) == 1 {
			space_switch = true
		} else if space_switch == true{
			array_count++
			space_switch = false
		}
	}
	return array_count
}

func SplitWhiteSpaces(s string) []string {
	array_count := Split_count(s)
	ret_array := make([]string, array_count)
	if array_count == 0 {
		return ret_array
	}
	i := 0
	for _,c := range s {
		//if c == ' ' {
		if isSpace(c) == 1 {
			i++
		} else {
			ret_array[i] += string(c)
		}
	}
	return ret_array
}