package piscine

func Join(strs []string, sep string) string {
	var word string
	for index, words := range strs {
		if index != 0 {
			word += sep
		} 
		word = word + words
	}
	return word
}
