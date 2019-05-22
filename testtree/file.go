package testtree

import "github.com/vektah/gqlparser/ast"

const (
	BeforeDirective = "before"
	AfterDirective  = "after"
)

type File struct {
	name   string
	before []*ast.OperationDefinition
	tests  []*ast.OperationDefinition
	after  []*ast.OperationDefinition
}

func findAndRemove(list *ast.DirectiveList, name string) *ast.Directive {
	for i, d := range *list {
		if d.Name == name {
			*list = append((*list)[:i], (*list)[i+1:]...)
			return d
		}
	}
	return nil
}

func NewFile(doc *ast.QueryDocument, name string) *File {
	f := &File{
		name: name,
	}

	for _, op := range doc.Operations {
		before := findAndRemove(&op.Directives, BeforeDirective)
		after := findAndRemove(&op.Directives, AfterDirective)

		if before != nil {
			f.before = append(f.before, op)
		}
		if after != nil {
			f.after = append(f.after, op)
		}
		if before == nil && after == nil {
			f.tests = append(f.tests, op)
		}
	}

	return f
}
