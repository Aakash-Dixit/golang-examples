package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	valMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		compliment := target - nums[i]
		if index, ok := valMap[compliment]; ok {
			result = append(result, index, i)
			break
		} else {
			valMap[nums[i]] = i
		}
	}
	return result
}

func main() {
	items := []int{2, 7, 11, 15}
	items = twoSum(items, 13)
	fmt.Println(items)
}
