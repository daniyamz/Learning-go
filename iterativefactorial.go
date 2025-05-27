/*Write an iterative function that returns the factorial of the int passed as parameter.

Errors (non possible values or overflows) will return 0.*/

package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= nb; i++ {
		if result > (1<<63-1)/i {
			return 0
		}
		result *= i
	}
	return result
}
