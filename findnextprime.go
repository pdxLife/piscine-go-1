package piscine

func FindNextPrime(nb int) int {
	for i := 2; i < nb; i++ {
		if nb % i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
