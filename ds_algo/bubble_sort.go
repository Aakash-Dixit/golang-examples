package main

import "fmt"

func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func main() {
	toBeSorted := []int{2, 10, 5, 7, 1, 20, 13}
	fmt.Println("before sorting : ", toBeSorted)
	bubbleSort(toBeSorted)
	fmt.Println("after sorting : ", toBeSorted)
}
