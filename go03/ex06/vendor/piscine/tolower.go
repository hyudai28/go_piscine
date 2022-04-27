package piscine

//import "fmt"

func ToLower(s string) string {
	lower_s := []rune(s)

	for i, c := range lower_s {
		if 'A' <= c && c <= 'Z' {
			lower_s[i] +=  32
		}
	}
	return (string)(lower_s)
}