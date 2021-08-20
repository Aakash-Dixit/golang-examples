package main

import (
	"fmt"
	"sort"
)

type programmer struct {
	Age int
}

type byAge []programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func main() {
	myInts := []int{7, 5, 1, 9, 2, 8}
	sort.Ints(myInts)
	fmt.Println("increasing order : ", myInts)
	sort.Sort(sort.Reverse(sort.IntSlice(myInts)))
	fmt.Println("decreasing order : ", myInts)

	programmers := []programmer{
		{Age: 30},
		{Age: 20},
		{Age: 50},
		{Age: 1000},
	}

	sort.Sort(byAge(programmers))

	fmt.Println(programmers)
}
