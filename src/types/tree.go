package types

type node[T any] struct {
	Value    T
	Children []*node[T]
}

type BinTree struct {
	root *node[int]
}

func (bt *BinTree) Add(value int) bool {
	if bt.root == nil {
		bt.root = &node[int]{Value: value}
		return true
	}
	return bt.add(value, bt.root)
}
func (bt *BinTree) add(value int, childNode *node[int]) bool {
	if value == childNode.Value {
		return false
	}
	if value < childNode.Value {
		if childNode.Children[0] == nil {
			childNode.Children[0] = &node[int]{Value: value}
			return true
		}
		return bt.add(value, childNode.Children[0])
	} else {
		if childNode.Children[1] == nil {
			childNode.Children[1] = &node[int]{Value: value}
			return true
		}
		return bt.add(value, childNode.Children[1])
	}
}

func (bt *BinTree) Get(value int) bool {
	return bt.get(value, bt.root)
}

func (bt *BinTree) get(value int, childNode *node[int]) bool {
	if childNode == nil {
		return false
	}
	if value == childNode.Value {
		return true
	}
	if value < childNode.Value {
		return bt.get(value, childNode.Children[0])
	}
	return bt.get(value, childNode.Children[1])
}
