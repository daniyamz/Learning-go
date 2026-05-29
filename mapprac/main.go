package main

import "fmt"

//Given an array of integers nums and an integer target, return the indices of the two numbers that add up to target.
//You may assume exactly one solution exists, and you may not use the same element twice.

func returnIdices(num []int, target int) []int {
	//numbs := make(map[int]int)
	var result []int
	for i := range num {
		for j := 1; j < len(num); j++ {
			if num[i]+num[j] == target {

				result = append(result, i, j)

			}
		}
	}
	return result
}

func main() {
	numb := []int{2, 13, 17, 7}
	res := returnIdices(numb, 9)
	fmt.Println(res) // [0, 1]
}
