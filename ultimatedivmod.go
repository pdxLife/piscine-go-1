package piscine

func UltimateDivMod(a *int, b *int) {
	var div int
	var mod int
	x := 13
	y := 2
	div = x / y
	mod = x % y
	*a = div
	*b = mod
}
