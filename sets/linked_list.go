package sets

import (
	"fmt"

	"github.com/lfmexi/datastructs/nodes"
)

// LinkedList is a dynamic list of linked set nodes and can contains any type of values
type LinkedList struct {
	first *nodes.DoubleLinkNode
	last  *nodes.DoubleLinkNode
	Size  int
}

// NewLinkedList returns a new instance of LinkedList
func NewLinkedList() LinkedList {
	return LinkedList{}
}

// Add creates a new element on the LinkedList
func (l *LinkedList) Add(v interface{}) interface{} {
	var newNode = nodes.NewDoubleLinkNode(v)
	if l.first == nil {
		l.first = newNode
		l.last = l.first
	} else {
		newNode.SetPrev(l.last)
		l.last.SetNext(newNode)
		l.last = newNode
	}
	l.Size++
	return v
}

// Delete removes the element located at i
func (l *LinkedList) Delete(i int) (interface{}, error) {
	if i >= l.Size || i < 0 {
		return nil, fmt.Errorf("Index %v out of bounds", i)
	}

	node := l.getNode(i)

	if node != nil {
		p, n := l.removeNode(node)
		if node == l.first {
			l.first = n
		}

		if node == l.last {
			l.last = p
		}

		return node.GetValue(), nil
	}

	return nil, nil
}

// Get returns the element that lives on the position i of the list
func (l *LinkedList) Get(i int) (interface{}, error) {
	if i >= l.Size || i < 0 {
		return nil, fmt.Errorf("Index %v out of bounds", i)
	}

	if i == l.Size-1 {
		return l.last.GetValue(), nil
	}

	node := l.getNode(i)

	if node != nil {
		return node.GetValue(), nil
	}

	return nil, nil
}

func (l *LinkedList) getNode(i int) *nodes.DoubleLinkNode {
	current := l.first
	for j := 0; current != nil && j <= i; j++ {
		if j == i {
			break
		}
		current = current.GetNext()
	}

	return current
}

// Pop returns the last inserted element, removing it from the list
func (l *LinkedList) Pop() interface{} {
	if l.last == nil {
		return nil
	}

	p, _ := l.removeNode(l.last)
	v := l.last.GetValue()

	if l.last == l.first {
		l.first = p
	}

	l.last = p

	return v
}

func (l *LinkedList) removeNode(node *nodes.DoubleLinkNode) (*nodes.DoubleLinkNode, *nodes.DoubleLinkNode) {
	p, n := node.UnlinkNode()

	if p != nil {
		p.SetNext(n)
	}

	if n != nil {
		n.SetPrev(p)
	}

	l.Size--
	return p, n
}
