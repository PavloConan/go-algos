package structures

type QNode struct {
	Value int
	Next  *QNode
}

type Queue struct {
	Length int
	Head   *QNode
	Tail   *QNode
}

func (q *Queue) Peek() int {
	if q.Head == nil {
		return 0
	} else {
		return q.Head.Value
	}
}
