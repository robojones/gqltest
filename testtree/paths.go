package testtree

// GetTestPaths finds all paths in the tree that resolve to a node containing tests.
func GetTestPaths(tree *Tree) []Path {
	return getTestPaths(nil, tree.root)
}

func getTestPaths(base Path, node *Node) []Path {
	var p []Path
	if node.IsTestNode() {
		p = append(p, base)
	}
	for name, child := range node.children {
		childBase := append(base, name)
		paths := getTestPaths(childBase, child)

		p = append(p, paths...)
	}
	return p
}
