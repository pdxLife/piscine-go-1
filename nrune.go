package piscine

func NRune(s string, n int) rune {
	strToRune := []rune(s)
	return strToRune[n-1]
}
