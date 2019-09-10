package linkedlist

import (
	"fmt"
)

type Node struct {
	Next *Node
	Key  interface{}
}

func (node *Node) Insert(key interface{}) {
	pointer := node
	for pointer.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = &Node{nil, key}
}

func (node *Node) Print() {
	if node == nil {
		return
	}

	if node.Next == nil {
		fmt.Printf("%v\n", node.Key)
		return
	}

	fmt.Printf("%v ->", node.Key)
	node.Next.Print()
}
