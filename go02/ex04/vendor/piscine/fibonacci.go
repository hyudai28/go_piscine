package piscine

func Fibonacci(index int) int {
	var ret int = 0

	if index < 0 {
		return -1
	} else if index == 1 {
		return 1
	} else if index == 0 {
		return 0
	}
	ret = Fibonacci(index - 1) + Fibonacci(index - 2)
	return ret
}
