package main

import "fmt"

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	x := make([]int, len(colsum))
	y := make([]int, len(colsum))
	sumX := 0
	sumY := 0
	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 2 {
			sumX++
			sumY++
			x[i] = 1
			y[i] = 1
		} else if colsum[i] == 1 {
			sumX++
			x[i] = 1
			y[i] = 0
		}
	}
	if sumX+sumY != upper+lower {
		return [][]int{}
	}
	if sumX > upper {
		for i := 0; i < len(x); i++ {
			if x[i] == 1 && y[i] == 0 {
				x[i], y[i] = y[i], x[i]
				sumX--
				sumY++
			}
			if sumX == upper {
				return [][]int{x, y}
			}
		}
	} else if sumX < upper {
		for i := 0; i < len(x); i++ {
			if x[i] == 0 && y[i] == 1 {
				x[i], y[i] = y[i], x[i]
				sumX++
				sumY--
			}
			if sumX == upper {
				return [][]int{x, y}
			}
		}
	}
	if sumX == upper {
		return [][]int{x, y}
	}
	return [][]int{}
}

func main() {
	fmt.Println(reconstructMatrix(5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}))
}
