package testtree

import "path"

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{
		root: NewNode(),
	}
}

func (t *Tree) SetFile(filepath string, f *File) {
	dir, name := path.Split(filepath)
	p := NewPath(dir)

	current := t.root

	for _, n := range p {
		next := current.children[n]
		if next == nil {
			next = AddChild(current, n)
		}
		current = next
	}

	current.files[name] = f
}
