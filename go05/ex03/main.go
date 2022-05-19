package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello\thow are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("hello"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(""))
}
