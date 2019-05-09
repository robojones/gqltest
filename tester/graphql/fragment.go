package graphql

import "github.com/vektah/gqlparser/ast"

func serializeFragment(frag *ast.FragmentDefinition, indent int) string {
	return fragment +
		frag.Name +
		serializeVariableList(frag.VariableDefinition) +
		on +
		frag.TypeCondition +
		serializeDirectiveList(frag.Directives) +
		serializeSelectionSet(frag.SelectionSet, indent)
}

func serializeFragmentSpread(frag *ast.FragmentSpread) string {
	return spread + frag.Name + serializeDirectiveList(frag.Directives)
}
