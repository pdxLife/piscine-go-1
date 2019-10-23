package piscine

func BasicJoin(strs []string) string {
	var word string
	for _, words := range strs {
		word += words 
	}
	return word
}
