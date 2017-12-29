package structures

import "errors"

type Queue struct {
	elements []*Vertex
	Size 	 int
	IsEmpty  bool
}

func NewQueue() Queue {
	return Queue{make([]*Vertex, 0), 0, true}
}

func (q *Queue) Front() (*Vertex, error) {
	if q.IsEmpty {
		return nil, errors.New("Enqueue is empty")
	}

	return q.elements[0], nil
}

func (q *Queue) Enqueue(element *Vertex) {
	q.elements = append(q.elements, element)
	q.Size++
	q.IsEmpty = false
}

func (q *Queue) Dequeue() (*Vertex, error) {

	element, err := q.Front()

	if err != nil {
		return nil, err
	}

	q.elements = q.elements[1:len(q.elements)]
	if len(q.elements) == 0 {
		q.IsEmpty = true
	}
	q.Size--

	return element, err
}