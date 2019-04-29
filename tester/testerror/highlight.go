package testerror

import (
	"fmt"
	"github.com/robojones/gqltest/source"
	"github.com/vektah/gqlparser/ast"
	"strings"
)

const NEWLINE = '\n'

func highlight(directive *ast.Position, operation *ast.Position) string {
	l := directive.Line

	runes := []rune(source.Content(operation))

	maxNumLen := getMaxLineNumberLength(runes, directive.Line)

	snippet := formatLineNumber(l, maxNumLen, directive.Line)

	for _, r := range runes {
		if r == NEWLINE {
			l += 1
			snippet += formatLineNumber(l, maxNumLen, directive.Line)
		}

		snippet += string(r)
	}

	return snippet + string(NEWLINE)
}

func formatLineNumber(l int, maxLen int, markLine int) string {
	prefix := fmt.Sprintf("%d", l)
	prefix = leftPad(prefix, maxLen)

	if l == markLine {
		prefix += "> "
	} else {
		prefix += "| "
	}

	return prefix
}

func getMaxLineNumberLength(runes []rune, startLine int) int {
	n := 0

	for _, r := range runes {
		if r == '\n' {
			n += 1
		}
	}

	return len(fmt.Sprintf("%d", startLine+n))
}

func leftPad(s string, n int) string {
	if len(s) > n {
		return s
	}
	return strings.Repeat(" ", n-len(s)) + s
}
