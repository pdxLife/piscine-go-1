package piscine

func AlphaCount(str string) int {
	counter := 0
	a := 0
	for range str {
		a++
	}
	for i := 0; i < a; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' {
			counter++
		}
	}
	return counter
}
