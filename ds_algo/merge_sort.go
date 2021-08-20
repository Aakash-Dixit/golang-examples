package main

import "fmt"

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items) / 2
	left := mergeSort(items[:mid])
	right := mergeSort(items[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var final []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	if len(left) > 0 {
		final = append(final, left[i:]...)
	}
	if len(right) > 0 {
		final = append(final, right[j:]...)
	}
	return final
}

func main() {
	toBeSorted := []int{2, 10, 5, 7, 1, 20, 13, 4}
	fmt.Println("before sorting : ", toBeSorted)
	sorted := mergeSort(toBeSorted)
	fmt.Println("after sorting : ", sorted)
}
