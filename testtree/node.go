package testtree

type Node struct {
	parent   *Node
	children map[string]*Node
	Files    map[string]*File
}

func NewNode() *Node {
	return &Node{
		children: make(map[string]*Node),
		Files:    make(map[string]*File),
	}
}

func AddChild(parent *Node, name string) (child *Node) {
	child = NewNode()
	child.parent = parent
	parent.children[name] = child
	return child
}

func (n *Node) IsTestNode() bool {
	for _, f := range n.Files {
		if len(f.tests) > 0 {
			return true
		}
	}
	return false
}
