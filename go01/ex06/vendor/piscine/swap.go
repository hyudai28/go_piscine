package piscine


func Swap(a, b *int) {
	var tmp int;

	tmp = *a;
	*a = *b;
	*b = tmp;
}