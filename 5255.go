package main

import "fmt"

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	for _, v := range indices {
		for i := 0; i < m; i++ {
			matrix[v[0]][i]++
		}
		for i := 0; i < n; i++ {
			matrix[i][v[1]]++
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j]%2 != 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println("leetcode solution")
	fmt.Println(oddCells(48, 37, [][]int{[]int{40, 5}}))
}
