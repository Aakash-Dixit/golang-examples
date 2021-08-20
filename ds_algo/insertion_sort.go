package main

import "fmt"

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func main() {
	toBeSorted := []int{2, 10, 5, 7, 1, 20, 13}
	fmt.Println("before sorting : ", toBeSorted)
	insertionSort(toBeSorted)
	fmt.Println("after sorting : ", toBeSorted)
}
