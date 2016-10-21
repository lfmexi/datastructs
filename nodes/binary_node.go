package nodes

// BinaryNode is a type of node that has two child nodes or links
// allows any type of value and index
type BinaryNode struct {
	index interface{}
	value interface{}
	prev  *BinaryNode
	next  *BinaryNode
}

// NewBinaryNode is a factory of BinaryNodes with values and indexes
func NewBinaryNode(index interface{}, value interface{}) *BinaryNode {
	return &BinaryNode{index, value, nil, nil}
}

// GetIndex gets the value that is stored on the node
func (b BinaryNode) GetIndex() interface{} {
	return b.index
}

// GetNext gets the next child node
func (b BinaryNode) GetNext() *BinaryNode {
	return b.next
}

// GetValue gets the value that is stored on the node
func (b BinaryNode) GetValue() interface{} {
	return b.value
}

// SetNext add a new BinaryNode as the next pointer
func (b *BinaryNode) SetNext(n *BinaryNode) {
	b.next = n
}

// SetIndex sets a new index to the node
func (b *BinaryNode) SetIndex(i interface{}) {
	b.index = i
}

// SetValue sets a new value to the node
func (b *BinaryNode) SetValue(v interface{}) {
	b.value = v
}
