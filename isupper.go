package piscine

func IsUpper(str string) bool {
	var t bool = true
	var f bool = false
	length := 0
	counter := 0
	for range str {
		length++
	}
	str_rune := []rune(str)
	for i := 0; i < length; i++ {
		if str_rune[i] >= 'A' && str_rune[i] <= 'Z' {
			counter++
		}

	}
	if length == counter {
		return t
	} else {
		return f
	}

}
