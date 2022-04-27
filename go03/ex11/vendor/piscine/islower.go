package piscine

func IsLower(s string) bool {
	for _, c := range s {
		if !('a' <= c && c <= 'z') {
			return false
		}
	}
	return true
}