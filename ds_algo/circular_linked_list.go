package main

import (
	"container/ring"
	"fmt"
)

func main() {
	myring := ring.New(5)

	l := myring.Len()
	for i := 0; i < l; i++ {
		myring.Value = i
		myring = myring.Next()
	}

	myring.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	fmt.Println("************* Next ************")
	// Iterate through the ring and print its contents
	for j := 0; j < l; j++ {
		fmt.Println(myring.Value)
		myring = myring.Next()
	}

	fmt.Println("************ Prev **************")
	// Iterate through the ring and print its contents
	myring = myring.Prev()
	for j := 0; j < l; j++ {
		fmt.Println(myring.Value)
		myring = myring.Prev()
	}

	myring1 := ring.New(3)

	for i := 5; i <= 7; i++ {
		myring1.Value = i
		myring1 = myring1.Next()
	}

	fmt.Println("******* Link *********")
	myring = myring.Link(myring1)
	myring.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	fmt.Println("******* Unink *********")
	myring.Unlink(3)
	myring.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}
