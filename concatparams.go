package piscine

func ConcatParams(args []string) string {
	l := 0
	result := ""
	for range args {
		l++
	}
	if l > 0 {
		for i := 0; i < l; i++ {
			result = result + args[i]
			if i < l-1 {
				result = result + "\n"
			}
		}
	}
	return result
}
