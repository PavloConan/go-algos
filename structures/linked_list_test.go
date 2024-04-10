// write tests for each linked linst function

package structures

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	if ll.Head.Value != 1 {
		t.Errorf("Expected %v, got %v", 1, ll.Head.Value)
	}

	if ll.Head.Next.Value != 2 {
		t.Errorf("Expected %v, got %v", 2, ll.Head.Next.Value)
	}

	if ll.Head.Next.Next.Value != 3 {
		t.Errorf("Expected %v, got %v", 3, ll.Head.Next.Next.Value)
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	ll := LinkedList{}

	ll.Prepend(1)
	ll.Prepend(2)
	ll.Prepend(3)

	if ll.Head.Value != 3 {
		t.Errorf("Expected %v, got %v", 3, ll.Head.Value)
	}

	if ll.Head.Next.Value != 2 {
		t.Errorf("Expected %v, got %v", 2, ll.Head.Next.Value)
	}

	if ll.Size != 3 {
		t.Errorf("Expected %v, got %v", 3, ll.Size)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.Remove(2)

	if ll.Head.Next.Value != 3 {
		t.Errorf("Expected %v, got %v", 3, ll.Head.Next.Value)
	}

	if ll.Size != 2 {
		t.Errorf("Expected %v, got %v", 2, ll.Size)
	}
}

func TestLinkedList_Get(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	node, err := ll.Get(1)
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if node.Value != 2 {
		t.Errorf("Expected %v, got %v", 2, node.Value)
	}

	_, err = ll.Get(3)
	if err == nil {
		t.Errorf("Expected an error, got %v", err)
	}
}

func TestLinkedList_InsertAt(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	err := ll.InsertAt(1, 4)
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if ll.Head.Next.Value != 4 {
		t.Errorf("Expected %v, got %v", 4, ll.Head.Next.Value)
	}

	err = ll.InsertAt(5, 5)
	if err == nil {
		t.Errorf("Expected an error, got %v", err)
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	err := ll.RemoveAt(1)
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if ll.Head.Next.Value != 3 {
		t.Errorf("Expected %v, got %v", 3, ll.Head.Next.Value)
	}

	err = ll.RemoveAt(5)
	if err == nil {
		t.Errorf("Expected an error, got %v", err)
	}
}

// function to test linked list search
func TestLinkedList_Search(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	if !ll.Search(2) {
		t.Errorf("Expected %v, got %v", true, false)
	}

	if ll.Search(5) {
		t.Errorf("Expected %v, got %v", false, true)
	}
}
