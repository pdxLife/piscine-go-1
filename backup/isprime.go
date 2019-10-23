package piscine

func IsPrime(nb int) bool {
	var t bool = true
	var f bool = false
	if nb == 0 || nb == 1 || nb == -1 {
		return f
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return f
		}
	}
	for i := -2; i > nb; i-- {
		if (nb%i) == 0 {
			return f
		}
	}
	return t
}
