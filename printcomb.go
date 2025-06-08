/*Write a function that prints, in ascending order and on a single line: all unique combinations of three
different digits so that, the first digit is lower than the second, and the second is lower than the third.

These combinations are separated by a comma and a space.*/

package piscine

import "github.com/01-edu/z01"

func PrintCombo() {
	for i := '0'; i <='9'; i++ {
		for j := i + 1; j <='9'; j++ {
			for k := j + 1; k <='9'; k ++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i !='7' || j !='8' || k != '9'{
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
