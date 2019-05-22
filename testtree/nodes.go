package testtree

type Nodes struct {
	path    Path
	index   int
	current *Node
}

// Returns an iterator which walks through the tree using the path.
func NewNodes(tree *Tree, path Path) *Nodes {
	return &Nodes{
		index:   0,
		path:    path,
		current: tree.root,
	}
}

func (n *Nodes) Value() *Node {
	return n.current
}

func (n *Nodes) Next() bool {
	if n.index >= len(n.path) {
		return false
	}
	name := n.path[n.index]

	child := n.current.children[name]
	if child == nil {
		return false
	} else {
		n.index += 1
		n.current = child
		return true
	}
}

func (n *Nodes) Previous() bool {
	parent := n.current.parent
	if parent == nil {
		return false
	} else {
		n.current = parent
		n.index -= 1
		return true
	}
}
