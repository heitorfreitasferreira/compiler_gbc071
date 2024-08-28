package types

type node[T comparable] struct {
	value    T
	children []*node[T]
}

type Tree[T comparable] struct {
	root *node[T]
}

func (t Tree[T]) Find(value T) (T, bool) {
	return t.find(t.root, value)
}

func (t Tree[T]) find(n *node[T], value T) (T, bool) {
	if n == nil {
		return value, false
	}
	if n.value == value {
		return value, true
	}
	for _, child := range n.children {
		if v, ok := t.find(child, value); ok {
			return v, true
		}
	}
	return value, false
}
