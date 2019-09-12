package binarysearchtree

type (
	Node struct {
		Key   int
		Left  *Node
		Right *Node
	}

	Tree struct {
		Root *Node
	}
)

func New() *Tree {
	return &Tree{}
}

func (t *Tree) SetRoot(n *Node) {
	t.Root = n
}

func (n *Node) insert(key int) {
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{key, nil, nil}
		} else {
			n.Left.insert(key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{key, nil, nil}
		} else {
			n.Right.insert(key)
		}
	}
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = &Node{key, nil, nil}
		return
	}
	t.Root.insert(key)
}
