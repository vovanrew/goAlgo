package structures

import "errors"

type Node struct {
	Key int
	Value int
	Next *Node
	Prev *Node
}

type List struct {
	Head *Node
}

func NewList() *List {
	return &List{nil}
}

func (l *List) AddNode(n *Node) {
	l.Head.Prev = n
	n.Next = l.Head
	l.Head = n
}

func (l *List) RemoveNodeWithKey(key int) {
	pointer := l.Head

	for pointer != nil {
		if pointer.Key == key {
			if pointer.Prev != nil {
				pointer.Prev.Next = pointer.Next
				if pointer.Next != nil {
					pointer.Next.Prev = pointer.Prev
				}
				break
			} else if pointer.Next != nil {
				pointer.Next.Prev = nil
			} else {
				l.Head = nil
			}
		}

		pointer = pointer.Next
	}
}

func (l *List) FindNodeWithKey(key int) (int, error) {
	pointer := l.Head

	for pointer != nil {
		if pointer.Key == key {
			return pointer.Value, nil
		}

		pointer = pointer.Next
	}

	return 0, errors.New("There is no element with such key")
}