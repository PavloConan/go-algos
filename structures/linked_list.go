package structures

import (
	"errors"
)

type LLNode struct {
	Value int
	Next  *LLNode
}

type LinkedList struct {
	Head *LLNode
	Size int
}
// 	for i := 0; i < 5000; i++ {
// 		unsortedArr[i] = i
// 	}

// 	unsortedArr.Shuffle()

func (ll *LinkedList) Append(val int) {
	newLLNode := &LLNode{
		Value: val,
	}

	if ll.Head == nil {
		ll.Head = newLLNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}

		current.Next = newLLNode
	}

	ll.Size++
}

func (ll *LinkedList) Prepend(val int) {
	newLLNode := &LLNode{
		Value: val,
	}

	if ll.Head == nil {
		ll.Head = newLLNode
	} else {
		currHead := ll.Head
		ll.Head = newLLNode
		ll.Head.Next = currHead
	}

	ll.Size++
}

func (ll *LinkedList) Remove(val int) {
	if ll.Size == 0 {
		return
	}

	if ll.Head.Value == val {
		ll.Head = ll.Head.Next
		ll.Size--
		return
	}

	curr := ll.Head

	for curr.Next != nil {
		if curr.Next.Value == val {
			break
		} else {
			curr = curr.Next
		}
	}

	curr.Next = curr.Next.Next
	ll.Size--
}

func (ll *LinkedList) Get(idx int) (*LLNode, error) {
	if idx >= ll.Size {
		return nil, errors.New("invalid index")
	}

	curr := ll.Head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}

	return curr, nil
}

func (ll *LinkedList) Search(val int) bool {
	curr := ll.Head

	for curr.Next != nil {
		if curr.Value == val {
			return true
		}
		curr = curr.Next
	}

	return false
}

func (ll *LinkedList) InsertAt(idx int, val int) error {
	if idx > ll.Size {
		return errors.New("invalid index")
	}

	if idx == ll.Size {
		ll.Append(val)
		return nil
	} else if idx == 0 {
		ll.Prepend(val)
		return nil
	}

	newNode := &LLNode{
		Value: val,
	}

	curr := ll.Head

	for i := 0; i < idx-1; i++ {
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode
	ll.Size++

	return nil
}

func (ll *LinkedList) RemoveAt(idx int) error {
	if ll.Head == nil {
		return nil
	}

	if idx >= ll.Size {
		return errors.New("invalid index")
	}

	curr := ll.Head

	for i := 0; i < idx-1; i++ {
		curr = curr.Next
	}

	removed := curr.Next
	curr.Next = removed.Next
	ll.Size--

	return nil
}

func (ll *LinkedList) Print() {
	if ll.Size == 0 {
		return
	}

	current := ll.Head
	for current != nil {
		if current.Next == nil {
			println(current.Value)
		} else {
			println(current.Value, "-> next:", current.Next.Value)
		}
		current = current.Next
	}
}
