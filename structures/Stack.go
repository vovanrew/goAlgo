package structures

import "errors"

type Stack struct {
	elements []*Vertex
	Size     int
	IsEmpty  bool
}

func NewStack() Stack {
	return Stack{make([]*Vertex, 0), 0, true}
}

func (s *Stack) Top() (*Vertex, error) {
	if s.IsEmpty {
		return nil, errors.New("Stack is empty")
	}

	return s.elements[len(s.elements) - 1], nil
}

func (s *Stack) Push(element *Vertex) {
	s.elements = append(s.elements, element)
	s.Size++
	s.IsEmpty = false
}

func (s *Stack) Pop() (*Vertex, error){
	element, err := s.Top()

	if err != nil {
		return nil, err
	}

	s.elements = s.elements[:len(s.elements) - 1]
	s.Size--
	if len(s.elements) == 0 {
		s.IsEmpty = true
	}
	return element, nil
}