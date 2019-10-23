package piscine

func ToLower(s string) string {
	length := len(s)
	s_rune := []rune(s)
	result := ""
	for i := 0; i < length; i++ {
		if s_rune[i] > 64 && s_rune[i] < 91 {
			s_rune[i] = s_rune[i] + 32
		}
		result = result + string(s_rune[i])
	}
	return result
}
