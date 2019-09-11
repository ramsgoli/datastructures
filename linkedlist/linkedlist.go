package linkedlist

import (
	"fmt"
)

type Node struct {
	Next *Node
	Key  interface{}
}

type LinkedList struct {
	Size int
	Head *Node
}

func New() *LinkedList {
	return &LinkedList{0, nil}
}

func (l *LinkedList) Insert(key interface{}) {
	if l.Head == nil {
		l.Head = &Node{nil, key}
		l.Size += 1
		return
	}

	pointer := l.Head
	for pointer.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = &Node{nil, key}

	l.Size += 1
}

func (l *LinkedList) InsertNode(newNode *Node) {
	if l.Head == nil {
		l.Head = newNode
		l.Size += 1
		return
	}

	pointer := l.Head
	for pointer.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = newNode

	l.Size += 1
}

// Remove removes a node with the specified value from the linked list
func (l *LinkedList) Remove(key interface{}) {
	if l.Head == nil {
		return
	}

	// case 1. The head node is the node to remove
	if l.Head.Key == key {
		l.Head = l.Head.Next
		l.Size -= 1
		return
	}
	// case 2.
	prev := l.Head
	curr := l.Head.Next

	// iterate through list until we find key we want to remove
	for curr != nil && curr.Key != key {
		prev = curr
		curr = curr.Next
	}

	if curr != nil {
		prev.Next = curr.Next
		l.Size -= 1
	}
}

func (l *LinkedList) Print() {
	if l.Head == nil {
		return
	}

	ptr := l.Head
	for ptr.Next != nil {
		fmt.Printf("%v -> ", ptr.Key)
		ptr = ptr.Next
	}
	fmt.Println(ptr.Key)
}
