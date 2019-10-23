package piscine

func AlphaCount(str string) []rune {
	//result := 0
	len := 0
	var res []rune
	var final []rune
	for i := range str {
		len = i
	}
	for i := 0; i <= len; i++ {
		if str[i] >= '0' && str[i] <= '9' || str[i] == '-' {
			res = append(res, rune(str[i]))
		}
	}
	for i := range res {
		if res[i] == '-' {
			if i == 0 {
				return res
			} else {
				for i := 0; i <= len; i++ {
					if str[i] >= '0' && str[i] <= '9' {
						final = append(final, rune(str[i]))
					}
				}
			}
			return final
		}
	}
	return res
}
