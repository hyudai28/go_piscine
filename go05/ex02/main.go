package main

import (
	"fmt"
	"piscine"
)

func main() {
	//test := []string{"Hello", "how", "are", "you?"}
	//test := []string{}
	//test := []string{""}
	test := []string{"Hello"}
	fmt.Println(piscine.ConcatParams(test))
}
