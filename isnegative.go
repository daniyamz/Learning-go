package piscine

import "github.com/01-edu/z01"

func IsNagetive(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
