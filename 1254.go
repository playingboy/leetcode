package main

import "fmt"

func closedIsland(grid [][]int) int {

	mark := make([][]int, len(grid))
	for index := 0; index < len(grid); index++ {
		mark[index] = make([]int, len(grid[1]))
	}
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if mark[i][j] == 0 && grid[i][j] == 0 {
				mark[i][j] = 1
				if Dfs(grid, mark, i, j) {
					ret++
				}
			}
		}
	}
	return ret
}

func Dfs(grid [][]int, mark [][]int, row, column int) bool {
	if row == 0 || row == len(grid)-1 || column == 0 || column == len(grid[0])-1 {
		return false
	}
	left := row - 1
	right := row + 1
	up := column + 1
	down := column - 1
	if grid[left][column] == 0 && mark[left][column] == 0 {
		mark[left][column] = 1
		if !Dfs(grid, mark, left, column) {
			return false
		}
	}
	if grid[right][column] == 0 {

	}
	
}

func main() {
	fmt.Println("leetcode solution")
	fmt.Println(closedIsland([][]int{[]int{1, 1, 1, 1, 1, 1, 1, 0},
		[]int{1, 0, 0, 0, 0, 1, 1, 0}, []int{1, 0, 1, 0, 1, 1, 1, 0}, []int{1, 0, 0, 0, 0, 1, 0, 1}, []int{1, 1, 1, 1, 1, 1, 1, 0}}))
}
