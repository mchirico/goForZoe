package main

import "fmt"

func populateDict(nums []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = 1
	}
	return m
}

func sorted(m map[int]int) []int {
	output := []int{}
	for i := 0; i < 100; i++ {
		if _, ok := m[i]; ok {
			output = append(output, i)
		}
	}
	return output
}

func main() {

	nums := []int{33, 51, 7, 11, 93}
	m := populateDict(nums)

	output := sorted(m)
	fmt.Printf("out: %v\n", output)
}

