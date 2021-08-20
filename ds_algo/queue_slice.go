package main

import "fmt"

//NodeQueue is a node in queue
type NodeQueue struct {
	Value int
}

// NewQueue returns a new queue with the given initial size.
func NewQueue() *Queue {
	return &Queue{}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*NodeQueue
	len   int
}

//Enqueue  adds a node to the queue.
func (q *Queue) Enqueue(value int) {
	n := &NodeQueue{value}
	q.nodes = append(q.nodes, n)
	q.len++
}

// Dequeue removes and returns a node from the queue in first to last order.
func (q *Queue) Dequeue() *NodeQueue {
	if q.len == 0 {
		return nil
	}
	dequqedNode := q.nodes[0]
	if q.len == 1 {
		q.nodes = q.nodes[:0]
	} else {
		q.nodes = q.nodes[1:]
	}
	q.len--
	return dequqedNode
}

// Peek returns the first element in the queue
func (q *Queue) Peek() *NodeQueue {
	return q.nodes[0]
}

// IsEmpty returns if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func main() {
	myqueue := NewQueue()
	fmt.Println("IsEmpty : ", myqueue.IsEmpty())

	myqueue.Enqueue(1)
	myqueue.Enqueue(2)
	myqueue.Enqueue(3)

	fmt.Println("len : ", myqueue.len)
	fmt.Println("dequeued : ", myqueue.Dequeue())
	fmt.Println("dequeued : ", myqueue.Dequeue())
	fmt.Println("dequeued : ", myqueue.Dequeue())
	fmt.Println("dequeued : ", myqueue.Dequeue())
	fmt.Println("len : ", myqueue.len)

	myqueue.Enqueue(4)
	myqueue.Enqueue(5)
	fmt.Println("len : ", myqueue.len)

	fmt.Println("peek : ", myqueue.Peek())

	fmt.Println("IsEmpty : ", myqueue.IsEmpty())
}
