package sets

import "testing"

func TestNewLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	if linkedList == nil {
		t.Error("The linked list has not been created")
	}
}

func TestAdd(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Add("element")
	if linkedList.Size != 1 {
		t.Error("The element has not been inserted")
	}
}

func TestDelete(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Add("element")
	linkedList.Add("Don't delete this one")
	element, err := linkedList.Delete(0)
	if err != nil {
		t.Error(err)
	}
	if element != "element" {
		t.Error("The element delete is not the right one")
	}
	if linkedList.Size > 1 {
		t.Error("The element hast not been deleted")
	}
}

func TestGet(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Add("element")
	linkedList.Add("Don't delete this one")
	element, err := linkedList.Get(0)
	if err != nil {
		t.Error(err)
	}
	if element != "element" {
		t.Error("The element is not the right one")
	}
}

func TestPop(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	element := linkedList.Pop()

	if element != 3 {
		t.Error("The element is not the right one")
	}

	if linkedList.Size != 2 {
		t.Error("The element has not been removed")
	}

	last, _ := linkedList.Get(linkedList.Size - 1)

	if last != 2 {
		t.Fail()
	}
}
