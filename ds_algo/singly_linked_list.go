package main

import (
	"errors"
	"fmt"
)

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

// Search returns node position with given value from linked list
func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
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

// Display displays all the nodes from linked list
func (l *LinkedList) Display() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

// Insert inserts new node at the end of  from linked list
func (l *LinkedList) Insert(values ...int) {
	for _, val := range values {
		n := Node{}
		n.value = val
		if l.len == 0 {
			l.head = &n
			l.tail = &n
			l.len++
			continue
		}
		l.tail.next = &n
		l.tail = l.tail.next
		l.len++
	}
}

// InsertAt inserts new node at given position
func (l *LinkedList) InsertAt(pos int, value int) {
	// create a new node
	newNode := Node{}
	newNode.value = value
	// validate the position
	if pos < 0 || pos > l.len {
		return
	} else if pos == 0 {
		newNode.next = l.head
		l.head = &newNode
		l.len++
		return
	} else if pos == l.len {
		l.tail.next = &newNode
		l.tail = &newNode
		l.len++
		return
	}
	ptr := l.head
	for p := 0; p < pos-1; p++ {
		ptr = ptr.next
	}
	newNode.next = ptr.next
	ptr.next = &newNode
	l.len++
	/*n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++*/
}

// DeleteAt deletes node at given position from linked list
func (l *LinkedList) DeleteAt(pos int) error {
	// validate the position
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	} else if pos < 0 || pos >= l.len {
		fmt.Println("position can not be negative or exceed linked lis length")
		return errors.New("position can not be negative or exceed linked lis length")
	} else if pos == 0 {
		l.head = l.head.next
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
	}
	/*prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos).next*/
	l.len--
	return nil
}

// DeleteVal deletes node having given value from linked list
func (l *LinkedList) DeleteVal(val int) error {
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	} else if l.head.value == val {
		l.head = l.head.next
		l.len--
		return nil
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next.value == val {
			ptr.next = ptr.next.next
			if ptr.next == l.tail {
				l.tail = ptr
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}

//Reverse reverses a linled list
func (l *LinkedList) Reverse() {
	var cur, prev, next *Node
	cur = l.head
	prev = nil
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = l.tail
}

func main() {
	l := LinkedList{}

	fmt.Println("\n************* Insert *************")
	//l.Insert(12)
	l.Insert(12, 13, 14)
	//l.Insert(14)
	l.Insert(15)
	fmt.Println("************* Print *************")
	l.Display()

	fmt.Println("\n************* InsertAt *************")
	l.InsertAt(4, 16)
	l.InsertAt(0, 11)
	l.InsertAt(-1, 13)
	l.InsertAt(1, 14)
	l.InsertAt(2, 15)
	fmt.Println("************* Print *************")
	l.Display()

	fmt.Println("\n************* Search *************")
	fmt.Println("Position of 12 value: ", l.Search(12))
	fmt.Println("Position of 14 value: ", l.Search(14))
	fmt.Println("Position of 15 value: ", l.Search(15))
	fmt.Println("Position of 100 value: ", l.Search(100))

	fmt.Println("\n************* GetAt *************")
	fmt.Println("Get At 1st position: ", l.GetAt(0))
	fmt.Println("Get At 3rd position: ", l.GetAt(2))
	fmt.Println("Get At 4th position: ", l.GetAt(3))
	fmt.Println("Get At -5 position (head is returned): ", l.GetAt(-5))

	fmt.Println("\n************* DeleteAt *************")
	l.DeleteAt(3)
	fmt.Println("************* Print *************")
	l.Display()

	fmt.Println("\n************* DeleteVal *************")
	l.DeleteVal(13)
	fmt.Println("************* Print *************")
	l.Display()

	fmt.Println("\n************* Reverse *************")
	l.Reverse()

	fmt.Println("************* Print *************")
	l.Display()
}
