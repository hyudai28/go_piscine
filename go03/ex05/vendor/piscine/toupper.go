package piscine

//import "fmt"

func ToUpper(s string) string {
	upper_s := []rune(s)

	for i, c := range upper_s {
		if 'a' <= c && c <= 'z' {
			upper_s[i] -=  32
		}
	}
	return (string)(upper_s)
}