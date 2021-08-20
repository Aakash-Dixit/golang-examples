package main

import "fmt"

type BinaryTree struct {
	root *BinaryNode
}

type BinaryNode struct {
	key   int
	left  *BinaryNode
	right *BinaryNode
}

// insert root in BinaryTree
func (t *BinaryTree) insert(data int) {
	if t.root == nil {
		t.root = &BinaryNode{key: data}
	} else {
		t.root.insert(data)
	}
}

// insert other nodes in BinaryTree
// insert only unique nodes
func (n *BinaryNode) insert(data int) {
	if data < n.key {
		if n.left == nil {
			n.left = &BinaryNode{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func (n *BinaryNode) search(key int) bool {
	if n == nil {
		return false
	}
	if n.key < key {
		return n.right.search(key)
	} else if n.key > key {
		return n.left.search(key)
	}
	return true
}

func (n *BinaryNode) FindMin() int {
	if n.left == nil {
		return n.key
	}
	return n.left.FindMin()
}

func (n *BinaryNode) FindMax() int {
	if n.right == nil {
		return n.key
	}
	return n.right.FindMax()
}

func (n *BinaryNode) delete(k int) *BinaryNode {
	if n == nil {
		return nil
	}
	if k < n.key {
		n.left = n.left.delete(k)
	} else if k > n.key {
		n.right = n.right.delete(k)
	} else {
		//hit the node to delete
		//Only one or zero child nodes
		if n.right == nil {
			return n.left
		}
		if n.left == nil {
			return n.right
		}
		// Have two child nodes at the same time

		// Get the inorder successor
		// (smallest in the right subtree)
		// Copy the inorder
		// successor's content to this node
		n.key = n.right.FindMin()

		// Delete the inorder successor
		n.right = n.right.delete(n.key)
	}
	return n
}

func printPreOrder(n *BinaryNode) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *BinaryNode) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%d ", n.key)
	}
}

func printInOrder(n *BinaryNode) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%d ", n.key)
		printInOrder(n.right)
	}
}

func main() {
	var t BinaryTree

	t.insert(1)
	t.insert(9)
	t.insert(4)
	t.insert(45)
	t.insert(31)
	t.insert(20)
	t.insert(100)
	t.insert(70)
	t.insert(81)

	fmt.Println("20 found :", t.root.search(20))
	fmt.Println("40 found :", t.root.search(40))

	fmt.Println("******* preorder *******")
	printPreOrder(t.root)

	fmt.Println()

	fmt.Println("******* postorder *******")
	printPostOrder(t.root)

	fmt.Println()

	fmt.Println("******* inorder *******")
	printInOrder(t.root)

	fmt.Println()

	fmt.Println("******* max *******")
	fmt.Println(t.root.FindMax())

	fmt.Println("******* min *******")
	fmt.Println(t.root.FindMin())

	fmt.Println("******* delete *******")
	t.root = t.root.delete(1)
	fmt.Println("root after deletion : ", t.root)

	fmt.Println("******* inorder *******")
	printInOrder(t.root)
}
