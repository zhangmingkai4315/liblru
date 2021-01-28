package liblru

import "errors"

var (
	ErrorLinkListEmpty = errors.New("list is empty")
)

type DoubleLinkList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkList()*DoubleLinkList{
	return &DoubleLinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (linklist *DoubleLinkList)Push(node *Node){
	if linklist.size == 0 {
		linklist.head = node
		linklist.tail = node
	}else{
		linklist.head.pre = node
		node.next = linklist.head
		node.pre = nil
		linklist.head = node
	}
	linklist.size++
}

func (linklist *DoubleLinkList) Pop() (*Node, error) {
	if linklist.size == 0 {
		return nil, ErrorLinkListEmpty
	}
	item := linklist.tail
	if linklist.size == 1{
		linklist.tail = nil
		linklist.head = nil
	}else{
		linklist.tail = linklist.tail.pre
		linklist.tail.next = nil
	}
	linklist.size --
	return item, nil
}
// Remove a node from double linklist
func (linklist *DoubleLinkList)Remove(node *Node){
	if linklist.size == 1{
		linklist.head = nil
		linklist.tail = nil
	}else{
		if node.pre != nil{
			node.pre.next = node.next
		}else{
			node.next.pre = nil
			linklist.head = node.next
		}
		if node.next != nil{
			node.next.pre = node.pre
		}else{
			node.pre.next = nil
			linklist.tail = node.pre
		}
	}
	linklist.size --
}

func (linklist *DoubleLinkList) Size()int {
	return linklist.size
}

func (linklist *DoubleLinkList) Tail()*Node {
	return linklist.tail
}

func (linklist *DoubleLinkList) Head()*Node {
	return linklist.head
}
// test helper function
func (linklist *DoubleLinkList) toList() []interface{} {
	current := linklist.head
	result := make([]interface{}, 0)
	for current != nil{
		result = append(result, current.value)
		current = current.next
	}
	return result
}