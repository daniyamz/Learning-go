/*Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.

A line is a sequence of characters preceding the end of line character ('\n').
*/
package main

import "github.com/01-edu/z01"

func main(){
	for k := '0' ; k <= '9' ; k ++{
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}