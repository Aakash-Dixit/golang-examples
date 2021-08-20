package main

import "fmt"

func binarySearch(key int, datalist []int) bool {
	low := 0
	high := len(datalist) - 1

	for low <= high {
		median := (low + high) / 2
		if datalist[median] == key {
			return true
		} else if datalist[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return false
}

func linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items)) //works on only sorted slices

	items = []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(linearsearch(items, 58))
}
