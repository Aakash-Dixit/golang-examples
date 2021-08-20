package main

import (
	"container/list"
	"fmt"
)

func main() {

	mylist := list.New()
	mylist.PushFront(1)
	mylist.PushBack(2)
	mylist.InsertBefore(0, mylist.Front())
	mylist.InsertAfter(3, mylist.Back())

	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}
	fmt.Println("************************************")
	mylist.MoveToFront(mylist.Front().Next())
	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

	fmt.Println("*************** print backward***************")
	for element := mylist.Back(); element != nil; element = element.Prev() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

	mylist.MoveToBack(mylist.Back().Prev())
	fmt.Println("************************************")
	for element := mylist.Back(); element != nil; element = element.Prev() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

	mylist.MoveAfter(mylist.Back().Prev(), mylist.Back())
	mylist.MoveBefore(mylist.Front().Next(), mylist.Front())
	fmt.Println("************************************")
	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

	fmt.Println("len of list : ", mylist.Len())

	fmt.Println("********** Remove *********")
	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		if element.Value == 0 {
			mylist.Remove(element)
		}
	}

	mylist.PushFrontList(mylist)
	mylist.PushBackList(mylist)
	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

}
