package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	array1 := getElements(root1)
	array2 := getElements(root2)
	result := make([]int, len(array1)+len(array2))
	l := 0
	i, j := 0, 0
	for i < len(array1) || j < len(array2) {
		if array1[i] < array2[j] {
			result[l] = array1[i]
		} else {
			result[l] = array2[j]
		}
		l++
	}
	if i < len(array1) {
		result[l] = array1[i]
		i++
		l++
	} else if j < len(array2) {
		result[l] = array1[j]
		j++
		l++
	}
	return result
}

func getElements(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	result = append(getElements(root.Left), result...)
	result = append(result, getElements(root.Right)...)
	return result
}

func main() {
	fmt.Println("leetcode solution")
	fmt.Println(getAllElements(nil, nil))
}
