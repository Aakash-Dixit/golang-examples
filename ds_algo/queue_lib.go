package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func main() {
	myqueue := queue.New()

	myqueue.Enqueue(1)
	myqueue.Enqueue(2)
	myqueue.Enqueue(3)
	myqueue.Enqueue(4)
	myqueue.Enqueue(5)

	fmt.Println("len : ", myqueue.Len())
	fmt.Println("peek : ", myqueue.Peek())

	fmt.Println("delete : ", myqueue.Dequeue())
	fmt.Println("delete : ", myqueue.Dequeue())

	fmt.Println("len : ", myqueue.Len())
	fmt.Println("peek : ", myqueue.Peek())
}
