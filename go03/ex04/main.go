package main
import (
        "fmt"
        "piscine"
)
func main() {
        fmt.Println(piscine.Compare("Hello!", "Hello!"))
        fmt.Println(piscine.Compare("Salut!", "lut!"))
        fmt.Println(piscine.Compare("Ola!", "Ol"))
        fmt.Println(piscine.Compare("Ola!", ""))
        fmt.Println(piscine.Compare("", "Ol"))
        fmt.Println(piscine.Compare("", ""))
}