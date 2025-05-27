/*Write a recursive function that returns the factorial of the int passed as parameter.

Errors (non possible values or overflows) will return 0.

for is forbidden for this exercise.
*/

package piscine

func RecursiveFactorial(nb int) int {

	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
