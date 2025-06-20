package leetcode

import (
	"log"
)

func LC146() {
	//TestCase1()
	//TestCase2()
	//TestCase4()
	TestCase7()
}

func TestCase4() {
	obj := Constructor(2)

	obj.Put(1, 1)
	//obj.Print()
	obj.Put(2, 2)
	//obj.Print()

	obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)

	//obj.Print()

	obj.Put(4, 4)
	//obj.Print()

	obj.Get(1)
	obj.Get(3)
	obj.Get(4)
}

func TestCase5() {
	obj := Constructor(1)
	obj.Put(2, 1)
	log.Println(obj.Get(2))
	obj.Put(3, 2)
	log.Println(obj.Get(2))
	log.Println(obj.Get(3))
}

func TestCase6() {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(2, 2)
	log.Println(obj.Get(2))
	obj.Put(1, 1)
	obj.Put(4, 1)
	log.Println(obj.Get(2))
}

func TestCase7() {
	obj := Constructor(1)
	obj.Put(2, 1)
	log.Println(obj.Get(2))
}

type LRUCache struct {
	capacity int
	cacheMap map[int]int
	nodeMap  map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	Prev *Node
	Key  int
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cacheMap: make(map[int]int, capacity),
		nodeMap:  make(map[int]*Node, capacity),
		capacity: capacity,
	}
}

func (lr *LRUCache) PopNode(key int) *Node {
	tmpNode := lr.nodeMap[key]
	if tmpNode.HasPrev() {
		prevNode := tmpNode.Prev
		if tmpNode.HasNext() {
			prevNode.Next = tmpNode.Next
		} else {
			prevNode.Next = nil
		}
	} else {
		lr.head = tmpNode.Next
	}

	if tmpNode.HasNext() {
		nextNode := tmpNode.Next
		if tmpNode.HasPrev() {
			nextNode.Prev = tmpNode.Prev
		} else {
			nextNode.Prev = nil
		}
	} else {
		lr.tail = tmpNode.Prev
	}

	tmpNode.Prev = nil
	tmpNode.Next = nil
	delete(lr.nodeMap, key)
	return tmpNode
}

func (lr *LRUCache) InsertTailNode(node *Node) {
	if lr.IsEmpty() {
		lr.head = node
		lr.tail = node
	} else {
		lr.tail.Next = node
		node.Prev = lr.tail
		lr.tail = node
	}
	lr.nodeMap[node.Key] = node
}

func (lr *LRUCache) Get(key int) int {
	if val, ok := lr.cacheMap[key]; ok {
		lr.ShiftRightNode(key)
		return val
	}
	return -1
}

func (lr *LRUCache) ShiftRightNode(key int) {
	poppedNode := lr.PopNode(key)
	tmpVal := lr.cacheMap[key]
	delete(lr.cacheMap, key)
	lr.InsertTailNode(poppedNode)
	lr.cacheMap[key] = tmpVal
}

func (lr *LRUCache) Put(key int, value int) {
	if _, ok := lr.cacheMap[key]; ok {
		lr.ShiftRightNode(key)
		lr.cacheMap[key] = value
		return
	}

	if lr.IsFull() {
		delete(lr.cacheMap, lr.head.Key)
		lr.PopNode(lr.head.Key)
	}

	newNode := &Node{Key: key}
	lr.InsertTailNode(newNode)
	lr.cacheMap[key] = value
}

func (lr *LRUCache) IsFull() bool {
	return len(lr.cacheMap) == lr.capacity
}

func (lr *LRUCache) IsEmpty() bool {
	return len(lr.cacheMap) == 0
}

func (n *Node) HasPrev() bool {
	return n.Prev != nil
}

func (n *Node) HasNext() bool {
	return n.Next != nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
