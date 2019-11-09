package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	// return 0 if input size is zero
	if len(nums) == 0 {
		return 0
	}

	// init table to count odds before [i]
	table := make([]int, len(nums)+1)
	count := 0 // count is the total subarrays
	m := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			table[i+1] = table[i] + 1
		} else {
			table[i+1] = table[i]
		}
		m[table[i+1]]++
		if table[i+1] >= k {
			count += m[table[i+1]-k]
		}
	}
	fmt.Println(m)
	return count
}

func main() {
	fmt.Println("leetcode solution")
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
