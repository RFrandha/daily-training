package other

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

type LinkList struct {
	head *Node
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

func (l *LinkList) AddHead(val int) {
	node := &Node{Val: val}
	if l.head == nil {
		l.head = node
	} else {
		node.Next = l.head
		l.head.Prev = node
	}
}

func (l *LinkList) RemoveHead() {
	l.head = l.head.Next
}

func (l *LinkList) RemoveIdx(idx int) {
	ptr := l.head
	for i := 0; i < idx-1; i++ {
		ptr = ptr.Next
	}
	ptr.Next = ptr.Next.Next
}

func LinkListTest() {

}
