package list

// LRUCache
// 146. LRU 缓存
// https://leetcode.cn/problems/lru-cache/
type LRUCache struct {
	len, cap   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	k, v       int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cap:   capacity,
		cache: make(map[int]*DLinkedNode, capacity),
		head:  &DLinkedNode{},
		tail:  &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}

	lru.moveToHead(node)
	return node.v
}

func (lru *LRUCache) Put(key int, value int) {
	exist, ok := lru.cache[key]
	if ok {
		exist.v = value
		lru.moveToHead(exist)
		return
	}

	node := &DLinkedNode{k: key, v: value}
	lru.cache[key] = node
	lru.addToHead(node)
	lru.len++
	if lru.len > lru.cap {
		removed := lru.removeTail()
		delete(lru.cache, removed.k)
		lru.len--
	}
}

func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *DLinkedNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
