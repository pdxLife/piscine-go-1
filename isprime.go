package piscine
	
func IsPrime(nb int) bool {
	var t bool = true
	var f bool = false
	for i := 2; i < nb; i++ {
		if nb % i == 0 {
			return f
		}
	}
	return t
}
