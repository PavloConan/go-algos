package structures

import "testing"

func TestDoublyLinkedList_Append(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)

	if dll.Head.Value != 1 {
		t.Errorf("Expected %v, got %v", 1, dll.Head.Value)
	}

	if dll.Head.Next.Value != 2 {
		t.Errorf("Expected %v, got %v", 2, dll.Head.Next.Value)
	}

	if dll.Head.Next.Prev.Value != 1 {
		t.Errorf("Expected %v, got %v", 1, dll.Head.Next.Prev.Value)
	}
}

func TestDoublyLinkedList_Prepend(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.Prepend(1)
	dll.Prepend(2)

	if dll.Head.Value != 2 {
		t.Errorf("Expected %v, got %v", 2, dll.Head.Value)
	}

	if dll.Head.Next.Value != 1 {
		t.Errorf("Expected %v, got %v", 1, dll.Head.Next.Value)
	}

	if dll.Head.Next.Prev.Value != 2 {
		t.Errorf("Expected %v, got %v", 2, dll.Head.Next.Prev.Value)
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	dll.Remove(2)

	if dll.Head.Next.Value != 3 {
		t.Errorf("Expected %v, got %v", 3, dll.Head.Next.Value)
	}

	if dll.Head.Next.Prev.Value != 1 {
		t.Errorf("Expected %v, got %v", 1, dll.Head.Next.Prev.Value)
	}
}

// function to test doubly linked list insert at index include errro check on invalid indx
func TestDoublyLinkedList_InsertAt(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	err := dll.InsertAt(1, 4)

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if dll.Head.Next.Value != 4 {
		t.Errorf("Expected %v, got %v", 4, dll.Head.Next.Value)
	}

	err = dll.InsertAt(5, 5)

	if err == nil {
		t.Errorf("Expected %v, got %v", "invalid index", err)
	}
}

// function to test double lniked list get functino with error handling
func TestDoublyLinkedList_Get(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	node, err := dll.Get(1)

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if node.Value != 2 {
		t.Errorf("Expected %v, got %v", 2, node.Value)
	}

	_, err = dll.Get(3)

	if err == nil {
		t.Errorf("Expected an error, got %v", err)
	}
}

// function to test doubly linked list search function
func TestDoublyLinkedList_Search(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	if !dll.Search(2) {
		t.Errorf("Expected %v, got %v", true, false)
	}

	if dll.Search(4) {
		t.Errorf("Expected %v, got %v", false, true)
	}
}

// function to test doubly linked list remove at index function with error handling
func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	err := dll.RemoveAt(1)

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if dll.Head.Next.Value != 3 {
		t.Errorf("Expected %v, got %v", 3, dll.Head.Next.Value)
	}

	err = dll.RemoveAt(5)

	if err == nil {
		t.Errorf("Expected %v, got %v", "invalid index", err)
	}
}
