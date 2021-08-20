package main

import "fmt"

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i + 1; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func main() {
	toBeSorted := []int{2, 10, 5, 7, 1, 20, 13}
	fmt.Println("before sorting : ", toBeSorted)
	selectionsort(toBeSorted)
	fmt.Println("after sorting : ", toBeSorted)
}
