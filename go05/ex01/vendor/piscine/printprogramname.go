package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	s := make([]int, max-min, max-min)
	for i, _ := range s {
		s[i] = min
		min++
	}
	return s
}
