package types

import (
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	Value    T
	Children []*Node[T]
}

func (n *Node[T]) AddNode(t T) {
	if n.Children == nil {
		n.Children = []*Node[T]{{
			Value:    t,
			Children: []*Node[T]{},
		}}
		return
	}

	n.Children = append(n.Children, &Node[T]{
		Value:    t,
		Children: []*Node[T]{},
	})
}

type Tree[T comparable] struct {
	Root *Node[T]
}

func (t Tree[T]) Find(value T) (T, bool) {
	return t.find(t.Root, value)
}

func (t Tree[T]) find(n *Node[T], value T) (T, bool) {
	if n == nil {
		return value, false
	}
	if n.Value == value {
		return value, true
	}
	for _, child := range n.Children {
		if v, ok := t.find(child, value); ok {
			return v, true
		}
	}
	return value, false
}

func (t Tree[T]) stringify(n *Node[T], depth int) string {
	if n == nil {
		return ""
	}
	result := fmt.Sprintf("%s%v\n", strings.Repeat("  ", depth), n.Value)
	for _, child := range n.Children {
		result += t.stringify(child, depth+1)
	}
	return result
}

func (t Tree[T]) String() string {
	return t.stringify(t.Root, 0)
}
