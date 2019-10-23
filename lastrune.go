package piscine

func LastRune(s string) rune {
	strToRune := []rune(s)
	n2 := -1
	for range s {
		n2++
	}
	return strToRune[n2]
}
