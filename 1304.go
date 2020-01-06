package main

import "fmt"

func sumZero(n int) []int {
	result := make([]int, 0)
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			result = append(result, 0)
		} else {
			result = append(result, -(j+1-i)/2)
			result = append(result, (j+1-i)/2)
		}
	}
	return result
}

func main() {
	fmt.Println("leetcode solution")
	fmt.Println(sumZero(100))
}
