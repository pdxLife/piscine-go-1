package piscine

func ToUpper(s string) string {
	length := len(s)
	s_rune := []rune(s)
	result := ""
	for i := 0; i < length; i++ {
		if s_rune[i] > 96 && s_rune[i] < 123 {
			s_rune[i] = s_rune[i] - 32
		}
		result = result + string(s_rune[i])
	}
	return result
}
