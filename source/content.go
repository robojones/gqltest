package source

import (
	"github.com/vektah/gqlparser/ast"
)

func Content(pos *ast.Position) string {
	runes := []rune(pos.Src.Input)

	return string(runes[pos.Start:pos.End])
}
