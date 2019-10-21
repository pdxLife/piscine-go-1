package piscine

func FindNextPrime(nb int) int {
	a := nb + 1
	for i := 0; i <= a; i++ {
		if i == nb || i == a {
			return i
		}
	}
	return 0
}
