package main

import "fmt"

func issnakeCamelCase(s string) bool {
	for _, cha := range s {
		if cha == '_' {
			return true
		}
	}
	return false
}

func SnakeCaseToCamelCase(s string) string {
	var result []rune
	UpperNext := false
	for _, cha := range s {
		if cha == '_' {
			UpperNext = true
		} else {
			if UpperNext {
				if cha >= 'a' && cha <= 'z' {
					cha -= 'a' - 'A'
					UpperNext = false
				}
			} 
			result = append(result, cha)
		}
	}
	return string(result)
}

func CamelToSnakeCase(s string) string {
	var result []rune
	for i, cha := range s {
		if cha >= 'A' && cha <= 'Z' {
			if i != 0 {
				result = append(result, '_')
			}
			cha += 'a' - 'A'
		}
		result = append(result, cha)
	}
	return string(result)
}

func main(){
	var input string
	fmt.Println("Enter a string (CamelCase or snake_case ):")
	fmt. Scanln(&input)

	if issnakeCamelCase(input){
		fmt.Println("Camel Case :", SnakeCaseToCamelCase(input))
	}else{
		fmt.Println("Snake Case :", CamelToSnakeCase(input))
	}

}
