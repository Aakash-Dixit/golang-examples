package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	myStack := stack.New()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println("stack len : ", myStack.Len())
	fmt.Println("stack peek : ", myStack.Peek())
	fmt.Println("popped val : ", myStack.Pop())

}
