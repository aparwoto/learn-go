package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (t *Node) insertLeft(data int) {
	t.left = &Node{data, nil, nil}
}

func (t *Node) insertRight(data int) {
	t.right = &Node{data: data} // another way to initialize struct both left and right will set automatically into nil
}

func printPreOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		printPreOrder(root.left)
		printPreOrder(root.right)
	}
}

func printInOrder(root *Node) {
	if root != nil {
		printInOrder(root.left)
		fmt.Printf("%d ", root.data)
		printInOrder(root.right)
	}
}

func printPostOrder(root *Node) {
	if root != nil {
		printPostOrder(root.left)
		printPostOrder(root.right)
		fmt.Printf("%d ", root.data)
	}
}

func main() {
	/*
	         1
	        / \
	       2   3
	      /   / \
	     4   5   6
	    /         \
	   7           8
	                \
	                 9
	*/
	root := &Node{1, nil, nil}
	root.insertLeft(2)
	root.insertRight(3)

	root.left.insertLeft(4)
	root.right.insertLeft(5)
	root.right.insertRight(6)

	root.left.left.insertLeft(7)
	root.right.right.insertRight(8)

	root.right.right.right.insertRight(9)

	fmt.Print("Print Pre Order: ")
	printPreOrder(root)
	println()
	fmt.Print("Print In Order: ")
	printInOrder(root)
	println()
	fmt.Print("Print Post Order: ")
	printPostOrder(root)

}
