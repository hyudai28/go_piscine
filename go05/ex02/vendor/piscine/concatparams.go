package piscine

import "fmt"

func ConcatParams(args []string) string {
	var ret_params string
	string_index := 0

	for range args {
		string_index++
	}
	for i, s := range args {
		ret_params += s
		if i < string_index - 1 {
			ret_params += "\n"
		}
	}
	return fmt.Sprint(ret_params)
}