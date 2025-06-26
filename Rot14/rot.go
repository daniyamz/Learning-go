/*Write a function rot14 that returns the string within the parameter transformed into a rot14 string.
Each letter will be replaced by the letter 14 spots ahead in the alphabetical order.
'z' becomes 'n' and 'Z' becomes 'N'. The case of the letter stays the same.
*/

package main
 
import "fmt"

func Rot14(s string) string {
	result := []rune {}

	for _, i := range s {
		if i>='a' && i<= 'z'{
			i = ((i - 'a' + 14) % 26) + 'a'
		}else if i >= 'A' && i <= 'Z'{
			i =((i - 'A' + 14) % 26) +'A'
		}
		result = append(result, i)
	}
	return string(result)
}

func main(){
	result := Rot14("abcdefghijklmnopqrstuvwxyz")
	for _, x := range result{
		fmt.Printf("%c ", x)
	}
	fmt.Println()
}