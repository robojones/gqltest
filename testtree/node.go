package testtree

type Node struct {
	parent   *Node
	children map[string]*Node
	files    map[string]*File
}

func NewNode() *Node {
	return &Node{
		children: make(map[string]*Node),
		files:    make(map[string]*File),
	}
}

func AddChild(parent *Node, name string) (child *Node) {
	child = NewNode()
	child.parent = parent
	parent.children[name] = child
	return child
}
