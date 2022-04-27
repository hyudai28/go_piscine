package main

import (
	"fmt"
	"piscine"
)

func main() {
	var prime bool = false
	//for i := 0; i < 2147483647 ;i++ {
	for i := 0; i < 5000 ;i++ {
		prime = piscine.IsPrime(i)
		if prime == true {
			fmt.Println(i)
		}
	}
}