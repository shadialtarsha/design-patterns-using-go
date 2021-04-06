package builder

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type htmlElement struct {
	name, text string
	elements   []htmlElement
}

func (e *htmlElement) string() string {
	return e.toString(0)
}

func (e *htmlElement) toString(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.toString(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.name))
	return sb.String()
}

type htmlBuilder struct {
	rootName string
	root     htmlElement
}

func newHTMLBuilder(rootName string) *htmlBuilder {
	return &htmlBuilder{
		rootName,
		htmlElement{
			name:     rootName,
			text:     "",
			elements: []htmlElement{},
		},
	}
}

func (b *htmlBuilder) addChild(childName, childText string) *htmlBuilder {
	e := htmlElement{
		name:     childName,
		text:     childText,
		elements: []htmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
	return b
}

func (b *htmlBuilder) string() string {
	return b.root.string()
}
