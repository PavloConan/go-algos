package structures

import (
	"errors"
	"fmt"
)

type DLLNode struct {
	Value int
	Next  *DLLNode
	Prev  *DLLNode
}

type DoublyLinkedList struct {
	Head *DLLNode
	Size int
}

func (dll *DoublyLinkedList) Append(val int) {
	newNode := &DLLNode{
		Value: val,
		Next:  nil,
		Prev:  nil,
	}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Size++
		return
	}

	curr := dll.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
	newNode.Prev = curr
	dll.Size++
}

func (dll *DoublyLinkedList) Prepend(val int) {
	newNode := &DLLNode{
		Value: val,
	}

	if dll.Head == nil {
		dll.Head = newNode
	} else {
		dll.Head.Prev = newNode
		newNode.Next = dll.Head
		dll.Head = newNode
	}

	dll.Size++
}

func (dll *DoublyLinkedList) Remove(val int) {
	if dll.Head == nil {
		return
	}

	if dll.Head.Value == val {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	} else {
		curr := dll.Head
		for curr.Next != nil {
			if curr.Value == val {
				break
			} else {
				curr = curr.Next
			}
		}

		if curr.Next == nil {
			return
		}

		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
	}

	dll.Size--
}

func (dll *DoublyLinkedList) InsertAt(idx int, val int) error {
	if idx > dll.Size {
		return errors.New("invalid index")
	}

	if idx == 0 {
		dll.Prepend(val)
		return nil
	}

	if idx == dll.Size {
		dll.Append(val)
		return nil
	}

	newNode := &DLLNode{
		Value: val,
	}

	curr := dll.Head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}

	newNode.Next = curr
	newNode.Prev = curr.Prev
	curr.Prev.Next = newNode
	curr.Prev = newNode
	dll.Size++

	return nil
}

func (dll *DoublyLinkedList) RemoveAt(idx int) error {
	if dll.Head == nil {
		return nil
	}

	if idx >= dll.Size {
		return errors.New("invalid index")
	}

	curr := dll.Head

	for i := 0; i < idx; i++ {
		curr = curr.Next
	}

	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	dll.Size--

	return nil
}

func (dll *DoublyLinkedList) Get(idx int) (*DLLNode, error) {
	if idx >= dll.Size {
		return nil, errors.New("invalid index")
	}

	curr := dll.Head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}

	return curr, nil
}

func (dll *DoublyLinkedList) Search(val int) bool {
	curr := dll.Head

	for curr.Next != nil {
		if curr.Value == val {
			return true
		} else {
			curr = curr.Next
		}
	}

	return false
}

func (dll *DoublyLinkedList) Print() {
	if dll.Size == 0 {
		return
	}

	current := dll.Head
	for current != nil {
		if current.Next == nil {
			fmt.Println(current.Prev.Value, "<- ", current.Value, "-> None")
		} else if current.Prev == nil {
			fmt.Println("None <- ", current.Value, "-> ", current.Next.Value)
		} else {
			fmt.Println(current.Prev.Value, "<- ", current.Value, "-> ", current.Next.Value)
		}
		current = current.Next
	}
}
