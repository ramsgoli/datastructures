package linkedlist

import (
	"fmt"
)

type Node struct {
	Next *Node
	Key  interface{}
}

func LinkedList(key interface{}) *Node {
	node := Node{nil, key}
	return &node
}

func (node *Node) Insert(key interface{}) {
	pointer := node
	for pointer.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = &Node{nil, key}
}

func (node *Node) InsertNode(newNode *Node) {
	pointer := node
	for pointer.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = newNode
}

// Remove removes a node with the specified value from the linked list
func (node *Node) Remove(key interface{}) {
	// case 1. The head node is the node to remove
	if node.Key == key {
		node = node.Next
		return
	}
	// case 2.
	prev := node
	curr := node.Next

	// iterate through list until we find key we want to remove
	for curr != nil && curr.Key != key {
		prev = curr
		curr = curr.Next
	}

	if curr != nil {
		prev.Next = curr.Next
	}
}

func (node *Node) Print() {
	if node == nil {
		return
	}

	if node.Next == nil {
		fmt.Printf("%v\n", node.Key)
		return
	}

	fmt.Printf("%v -> ", node.Key)
	node.Next.Print()
}
