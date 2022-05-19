package piscine

func AppendRange(min, max int) []int {
	s := []int{}

	if min >= max {
		return nil
	}
	for min < max {
		s = append(s, min)
		min++
	}
	return s
}
