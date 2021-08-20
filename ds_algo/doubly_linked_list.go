package main

import (
	"errors"
	"fmt"
)

//NodeDoubly represents a node in a doubly linked list
type NodeDoubly struct {
	val  int
	prev *NodeDoubly
	next *NodeDoubly
}

//List represents a doubly linked list
type List struct {
	len  int
	head *NodeDoubly
	tail *NodeDoubly
}

// Insert inserts new node at the end of linked list
func (l *List) Insert(values ...int) {
	for _, val := range values {
		n := NodeDoubly{}
		n.val = val

		if l.len == 0 {
			l.head = &n
			l.tail = &n
			l.len++
			continue
		}
		l.tail.next = &n
		l.tail.next.prev = l.tail
		l.tail = l.tail.next
		l.len++
	}
}

//InsertAt inserts a node at position = pos
func (l *List) InsertAt(pos, val int) {
	if pos < 0 || pos > l.len {
		fmt.Println("cannot insert at pos : ", pos)
		return
	}
	n := NodeDoubly{}
	n.val = val
	if pos == 0 {
		head := l.head
		l.head.prev = &n
		l.head = &n
		l.head.next = head
		l.len++
		return
	}
	if pos == l.len {
		tail := l.tail
		l.tail.next = &n
		l.tail = &n
		l.tail.prev = tail
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	prev := ptr.prev
	prev.next = &n
	ptr.prev = &n
	ptr.prev.prev = prev
	ptr.prev.next = ptr
	l.len++
}

// Display lists the node of linked list
func (l *List) Display() {
	ptr := l.head
	for ptr != nil {
		fmt.Println(ptr)
		ptr = ptr.next
	}
}

// Search returns node position with given value from linked list
func (l *List) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.val == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// GetAt returns node at given position from linked list
func (l *List) GetAt(pos int) *NodeDoubly {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

// DeleteAt deletes node at given position from linked list
func (l *List) DeleteAt(pos int) error {
	// validate the position
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	} else if pos < 0 || pos >= l.len {
		fmt.Println("position can not be negative or exceed linked lis length")
		return errors.New("position can not be negative or exceed linked lis length")
	} else if pos == 0 {
		l.head = l.head.next
		l.head.prev = nil
		l.len--
		return nil
	}
	ptr := l.head
	for i := 0; i < pos-1; i++ {
		ptr = ptr.next
	}
	ptr.next = ptr.next.next
	if pos == l.len-1 {
		l.tail = ptr
	} else {
		ptr.next.prev = ptr
	}
	l.len--
	return nil
}

// DeleteVal deletes node having given value from linked list
func (l *List) DeleteVal(val int) error {
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	} else if l.head.val == val {
		l.head = l.head.next
		l.head.prev = nil
		l.len--
		return nil
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next.val == val {
			if ptr.next == l.tail {
				l.tail = ptr
				ptr.next = nil
			} else {
				ptr.next = ptr.next.next
				ptr.next.prev = ptr
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}

// Reverse reverses a linked lost
func (l *List) Reverse() {
	var cur, next *NodeDoubly
	cur = l.tail
	l.tail = l.head
	l.head = cur
	next = nil
	for cur != nil {
		next = cur.prev
		cur.prev = cur.next
		cur.next = next
		cur = next
	}
}

func main() {
	l := List{}

	fmt.Println("******* Insert ********")
	l.Insert(0, 1, 2, 3)
	l.Insert(4)

	fmt.Println("******** Display ********")
	l.Display()

	fmt.Println("******* Insert ********")
	l.InsertAt(5, 5)
	l.InsertAt(6, 6)
	l.InsertAt(7, 7)
	l.InsertAt(8, 8)
	l.InsertAt(9, 9)

	fmt.Println("******** Display ********")
	l.Display()

	fmt.Println("******* Search ********")
	fmt.Println("2 found at : ", l.Search(2))
	fmt.Println("5 found at : ", l.Search(5))

	fmt.Println("******* GetAt ********")
	fmt.Println("node at pos 0 : ", l.GetAt(0))
	fmt.Println("node at pos 3 : ", l.GetAt(3))
	fmt.Println("node at pos 6 : ", l.GetAt(6))

	fmt.Println("******* DeleteVal ********")
	l.DeleteVal(0)
	l.DeleteVal(5)
	l.DeleteVal(9)
	l.Display()

	fmt.Println("******* DeleteAt ********")
	l.DeleteAt(0)
	l.DeleteAt(2)
	l.DeleteAt(4)
	l.Display()

	fmt.Println("******* Reverse ********")
	l.Reverse()
	l.Display()
}
