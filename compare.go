package piscine

func Compare(a, b string) int {
	lena := 0
	//lenb := 0
	a_rune := []rune(a)
	b_rune := []rune(b)
	for range a_rune {
		lena++
	}
	if a == b {
		return 0
	} else if a_rune[0] == b_rune[0] && a_rune[1] == b_rune[1] {
		return 1
	} else {
		return -1
	}

}
