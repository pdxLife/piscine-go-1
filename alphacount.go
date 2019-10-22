package piscine

func AlphaCount(str string) int {
	counter := 0
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z'{
			counter++
		}
	}
	return counter
}
