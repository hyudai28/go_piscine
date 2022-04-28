package main
import (
"fmt"
"piscine"
)
func main() {
s := "Hello World!"
s = piscine.StrRev(s)
fmt.Println(s)
s = piscine.StrRev("")
fmt.Println(s)
}