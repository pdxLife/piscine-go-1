package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a1 := 48; a1 <= 57; a1++ {
		for a2 := 48; a2 <= 57; a2++ {
			for b1 := 48; b1 <= 57; b1++ {
				for b2 := 48; b2 <= 57; b2++ {
					if a1 != b1 || a2 != b2 {
						z01.PrintRune(rune(a1))
						z01.PrintRune(rune(a2))
						z01.PrintRune(32)
						z01.PrintRune(rune(b1))
						z01.PrintRune(rune(b2))
							if a1 != 57 || a2 != 57 || b1 != 57 || b2 != 56 {
								z01.PrintRune(44)
								z01.PrintRune(32)
							}
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
