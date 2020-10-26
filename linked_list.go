package linkedlist

//ErrEmptyList ...
var ErrEmptyList error

//Node ....
type Node struct {
	Val      int
	NextNode *Node
	PrevNode *Node
}

//List ...
type List struct {
	Head *Node
	Tail *Node
}

//NewList ...
func NewList(args ...interface{}) *List {
	if len(args) == 0 {
		return &List{nil, nil}
	}
	if len(args) == 1 {
		newNode := &Node{args[0].(int), nil, nil}
		return &List{newNode, newNode}
	}
	res := &List{}
	head := Node{}
	prev := &Node{}

	curr := &head
	for _, val := range args {

		curr.Val = val.(int)
		curr.PrevNode = prev
		prev.NextNode = curr
		prev = curr
		curr = &Node{}

	}

	prev.NextNode = nil
	headprev := head.PrevNode
	headprev.NextNode = nil
	head.PrevNode = nil

	res.Head = &head
	res.Tail = prev
	return res

}

//Next ...
func (e *Node) Next() *Node {
	if e != nil {
		return e.NextNode
	}
	return nil

}

//Prev ...
func (e *Node) Prev() *Node {
	if e != nil {
		return e.PrevNode
	}
	return nil

}

//PushFront ....
func (l *List) PushFront(v interface{}) {
	newNode := &Node{v.(int), nil, nil}
	if l.Head != nil {
		newNode.NextNode = l.Head
		l.Head.PrevNode = newNode
		l.Head = newNode
		return

	}
	l.Tail = newNode
	l.Head = newNode
	return

}

//PushBack ...
func (l *List) PushBack(v interface{}) {
	newNode := &Node{v.(int), nil, nil}
	if l.Tail != nil {
		newNode.PrevNode = l.Tail
		l.Tail.NextNode = newNode
		l.Tail = newNode
		return

	}
	l.Head = newNode
	l.Tail = newNode
	return

}

//PopFront ....
func (l *List) PopFront() (interface{}, error) {
	var resh interface{}
	if l.Head == nil {
		return nil, ErrEmptyList
	} else if l.Head.NextNode != nil {
		newHead := l.Head.NextNode
		newHead.PrevNode = nil
		l.Head.NextNode = nil

		resh = l.Head.Val
		l.Head = newHead
		return resh, nil

	}
	resh = l.Head.Val
	l.Head = nil
	l.Tail = nil
	return resh, nil

}

//PopBack ....
func (l *List) PopBack() (interface{}, error) {

	var rest interface{}
	if l.Tail == nil {
		return nil, ErrEmptyList
	} else if l.Tail.PrevNode != nil {
		newTail := l.Tail.PrevNode
		newTail.NextNode = nil
		l.Tail.PrevNode = nil

		rest = l.Tail.Val
		l.Tail = newTail
		return rest, nil

	}
	rest = l.Tail.Val
	l.Tail = nil
	l.Head = nil
	return rest, nil

}

//Reverse ....
func (l *List) Reverse() *List {
	curr := l.Head
	var p interface{}

	prev, _ := p.(*Node)
	for curr != nil {
		prev = curr.PrevNode
		curr.PrevNode = curr.NextNode
		curr.NextNode = prev
		curr = curr.PrevNode

	}
	if prev != (*Node)(nil) {
		l.Tail = l.Head
		l.Head = prev.PrevNode

		return l
	}
	return l

}

//First ....
func (l *List) First() *Node {
	return l.Head

}

//Last ...
func (l *List) Last() *Node {

	return l.Tail
}
