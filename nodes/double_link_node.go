package nodes

// DoubleLinkNode is a type of node that has two child nodes or links
// allows any type of value and index
type DoubleLinkNode struct {
	value interface{}
	prev  *DoubleLinkNode
	next  *DoubleLinkNode
}

// NewDoubleLinkNode is a factory of DoubleLinkNodes with values
func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value, nil, nil}
}

// GetNext gets the next child node
func (b DoubleLinkNode) GetNext() *DoubleLinkNode {
	return b.next
}

// GetPrev gets the previous child node
func (b DoubleLinkNode) GetPrev() *DoubleLinkNode {
	return b.prev
}

// GetValue gets the value that is stored on the node
func (b DoubleLinkNode) GetValue() interface{} {
	return b.value
}

// SetNext add a new DoubleLinkNode as the next pointer
func (b *DoubleLinkNode) SetNext(n *DoubleLinkNode) {
	b.next = n
}

// SetPrev add a new DoubleLinkNode as the previous pointer
func (b *DoubleLinkNode) SetPrev(n *DoubleLinkNode) {
	b.prev = n
}

// SetValue sets a new value to the node
func (b *DoubleLinkNode) SetValue(v interface{}) {
	b.value = v
}

// UnlinkNode removes the links of the related nodes
// returning prev,next pair
func (b *DoubleLinkNode) UnlinkNode() (*DoubleLinkNode, *DoubleLinkNode) {
	p := b.prev
	n := b.next
	b.prev, b.next = nil, nil
	return p, n
}
