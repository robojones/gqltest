package testerror

import (
	"fmt"
	"github.com/vektah/gqlparser/ast"
	"strings"
)

const (
	newline     = '\n'
	linesBefore = 1
	linesAfter  = 1
)

func highlight(directive *ast.Directive) string {
	pos := directive.Position
	selectIndex := pos.Line - 1

	lines := getLines(pos.Src)
	fromLine := selectIndex - linesBefore
	toLine := selectIndex + 1 + linesAfter
	maxNumberLen := getNumberLength(toLine)

	selectedLines := lines[fromLine:toLine]

	r := ""

	for i, line := range selectedLines {
		lineIndex := fromLine + i
		lineNumber := lineIndex + 1

		mark := lineIndex == selectIndex
		prefix := formatLineNumber(lineNumber, maxNumberLen, mark)

		line = prefix + line

		isLast := lineIndex == (selectIndex + linesAfter)
		if !isLast {
			line += string(newline)
		}

		r += line
	}

	return r
}

func getLines(src *ast.Source) []string {
	var runes = []rune(src.Input)

	var lineNumber = 0
	var lines = []string{""}

	for _, r := range runes {
		if r == newline {
			lines = append(lines, "")
			lineNumber += 1
		} else {
			lines[lineNumber] = lines[lineNumber] + string(r)
		}
	}

	return lines
}

func formatLineNumber(l int, maxLen int, markLine bool) string {
	prefix := fmt.Sprintf("%d", l)
	prefix = leftPad(prefix, maxLen)

	if markLine {
		prefix += "> "
	} else {
		prefix += "| "
	}

	return prefix
}

func getNumberLength(number int) int {
	return len(fmt.Sprintf("%d", number))
}

func leftPad(s string, n int) string {
	if len(s) > n {
		return s
	}
	return strings.Repeat(" ", n-len(s)) + s
}
