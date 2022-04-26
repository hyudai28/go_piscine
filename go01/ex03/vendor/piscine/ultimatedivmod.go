package piscine

func UltimateDivMod(a *int, b *int) {
	var div int
	var per int

	div = *a / *b
	per = *a % *b
	*a = div
	*b = per
}