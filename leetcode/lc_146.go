package leetcode

func LC146OLD() {
	//TestCase1()
	//TestCase2()
	//TestCase3()
}

func TestCase1() {
	//obj := Constructor(2)
	//
	//obj.Put(1, 1)
	////obj.Print()
	//obj.Put(2, 2)
	////obj.Print()
	//
	//obj.Get(1)
	//obj.Put(3, 3)
	//obj.Get(2)
	//
	////obj.Print()
	//
	//obj.Put(4, 4)
	////obj.Print()
	//
	//obj.Get(1)
	//obj.Get(3)
	//obj.Get(4)
}

func TestCase2() {
	//obj := Constructor(1)
	//obj.Put(2, 1)
	//log.Println(obj.Get(2))
	//obj.Put(3, 2)
	//log.Println(obj.Get(2))
	//log.Println(obj.Get(3))
}

func TestCase3() {
	//obj := Constructor(2)
	//obj.Put(2, 1)
	//obj.Put(2, 2)
	//log.Println(obj.Get(2))
	//obj.Put(1, 1)
	//obj.Put(4, 1)
	//log.Println(obj.Get(2))
}

//
//type LRUCache struct {
//	capacity int
//	cacheMap map[int]int
//	head     *Node
//	tail     *Node
//}
//type Node struct {
//	Prev *Node
//	Key  int
//	Next *Node
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		cacheMap: make(map[int]int, capacity),
//		capacity: capacity,
//	}
//}
//
//func (lr *LRUCache) Print() {
//	newTail := lr.tail
//	for newTail != nil {
//		fmt.Print(newTail.Key, " ")
//		newTail = newTail.Prev
//	}
//	fmt.Println()
//}
//
//func (lr *LRUCache) Get(key int) int {
//	if _, ok := lr.cacheMap[key]; ok {
//		newTail := lr.tail
//		for newTail != nil {
//			if newTail.Key == key {
//				ptrPrev := newTail.Prev
//				ptrNext := newTail.Next
//
//				if ptrPrev == nil && ptrNext == nil {
//					return lr.cacheMap[key]
//				}
//
//				if ptrPrev != nil {
//					ptrPrev.Next = newTail.Next
//				} else if ptrNext != nil {
//					lr.head = ptrNext
//				}
//				if ptrNext != nil {
//					ptrNext.Prev = ptrPrev
//				}
//
//				newTail.Next = nil
//				lr.insertTail(newTail)
//				return lr.cacheMap[key]
//			}
//			newTail = newTail.Prev
//		}
//	}
//	return -1
//}
//
//func (lr *LRUCache) Put(key int, value int) {
//	if _, ok := lr.cacheMap[key]; ok {
//		lr.cacheMap[key] = value
//		newTail := lr.tail
//		for newTail != nil {
//			if newTail.Key == key {
//				ptrPrev := newTail.Prev
//				ptrNext := newTail.Next
//				if ptrPrev != nil {
//					ptrPrev.Next = newTail.Next
//				} else if ptrNext != nil {
//					lr.head = ptrNext
//				}
//				if ptrNext != nil {
//					ptrNext.Prev = ptrPrev
//				}
//				break
//			}
//			newTail = newTail.Prev
//		}
//
//		newTail.Next = nil
//		lr.insertTail(newTail)
//	} else {
//		if len(lr.cacheMap) == 0 {
//			newNode := &Node{Key: key}
//			if lr.head == nil && lr.tail == nil {
//				lr.head = newNode
//				lr.tail = newNode
//			} else {
//				lr.tail.Next = newNode
//				newNode.Prev = lr.tail
//				lr.tail = newNode
//			}
//			lr.cacheMap[key] = value
//		} else if len(lr.cacheMap) == lr.capacity {
//			delete(lr.cacheMap, lr.head.Key)
//			if lr.head.Next != nil {
//				lr.head = lr.head.Next
//				lr.head.Prev = nil
//			}
//
//			lr.cacheMap[key] = value
//			newNode := &Node{Key: key}
//			lr.insertTail(newNode)
//		} else {
//			lr.cacheMap[key] = value
//			newNode := &Node{Key: key}
//			lr.insertTail(newNode)
//		}
//	}
//}
//
//func (lr *LRUCache) insertTail(newNode *Node) {
//	newNode.Prev = lr.tail
//	lr.tail.Next = newNode
//	lr.tail = newNode
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := ConstructorPoN(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
