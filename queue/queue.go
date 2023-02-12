package queue

// MyCircularDeque
// 641. 设计循环双端队列
// https://leetcode.cn/problems/design-circular-deque/
type MyCircularDeque struct {
	head, tail *node
	len, cap   int
}

type node struct {
	prev, next *node
	val        int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{cap: k}
}

func (deque *MyCircularDeque) InsertFront(value int) bool {
	if deque.IsFull() {
		return false
	}

	element := &node{val: value}
	if deque.IsEmpty() {
		deque.head = element
		deque.tail = element
		deque.head.next = element
		deque.head.prev = element
		deque.len++
		return true
	}

	element.next = deque.head
	deque.head.prev = element
	element.prev = deque.tail
	deque.head = element
	deque.len++
	return true
}

func (deque *MyCircularDeque) InsertLast(value int) bool {
	if deque.IsFull() {
		return false
	}

	element := &node{val: value}
	if deque.IsEmpty() {
		deque.head = element
		deque.tail = element
		deque.head.next = element
		deque.head.prev = element
		deque.len++
		return true
	}

	element.prev = deque.tail
	deque.tail.next = element
	element.next = deque.head
	deque.tail = element
	deque.len++
	return true
}

func (deque *MyCircularDeque) DeleteFront() bool {
	if deque.len == 0 {
		return false
	}

	deque.len--
	if deque.len == 0 {
		deque.head = nil
		deque.tail = nil
		return true
	}

	oldHead := deque.head
	deque.head = oldHead.next
	oldHead.next = nil
	deque.head.prev = deque.tail
	return true
}

func (deque *MyCircularDeque) DeleteLast() bool {
	if deque.len == 0 {
		return false
	}

	deque.len--
	if deque.len == 0 {
		deque.head = nil
		deque.tail = nil
		return true
	}

	oldTail := deque.tail
	deque.tail = oldTail.prev
	oldTail.next = nil
	deque.tail.next = deque.head
	return true
}

func (deque *MyCircularDeque) GetFront() int {
	if deque.IsEmpty() {
		return -1
	}
	return deque.head.val
}

func (deque *MyCircularDeque) GetRear() int {
	if deque.IsEmpty() {
		return -1
	}
	return deque.tail.val
}

func (deque *MyCircularDeque) IsEmpty() bool {
	return deque.len == 0
}

func (deque *MyCircularDeque) IsFull() bool {
	return deque.len == deque.cap
}
