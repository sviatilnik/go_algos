package liskenlist

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head   *Node[T]
	Length int
}

func (l *LinkedList[T]) Insert(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Length++
		return
	}

	node := l.Head
	for node.Next != nil {
		node = node.Next
	}

	node.Next = newNode
	l.Length++
}

func (l *LinkedList[T]) Remove(value T) {
	if l.Head == nil {
		return
	}

	node := l.Head
	var prev *Node[T]
	for node.Value != value {
		prev = node
		node = node.Next
	}

	if prev != nil {
		prev.Next = node.Next
		node.Next = nil
		l.Length--
		return
	}

	// head
	l.Head = nil
	l.Length = 0
}

func (l *LinkedList[T]) String() string {
	str := ""

	node := l.Head
	for node != nil {
		str = str + fmt.Sprintf("%v", node.Value) + " -> "
		node = node.Next
	}

	return str
}

func LinkedListSample() {
	list := new(LinkedList[int])

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	fmt.Println(list)

	list.Remove(2)
	fmt.Println(list)
	list.Remove(1)
	fmt.Println(list)

	list.Insert(3)
	fmt.Println(list)
}
