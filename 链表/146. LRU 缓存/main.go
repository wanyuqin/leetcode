package main

import "container/list"

type LRUCache struct {
	cap int
	len int

	list *list.List

	keyToNode map[int]*list.Element
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:       capacity,
		list:      list.New(),
		keyToNode: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.keyToNode[key]
	if node == nil { // 没有这本书
		return -1
	}
	this.list.MoveToFront(node) // 把这本书放在最上面
	return node.Value.(entry).value

}

func (this *LRUCache) Put(key int, value int) {
	if node := this.keyToNode[key]; node != nil {
		// 节点存在
		node.Value = entry{key: key, value: value}
		// 把这个节点放在最前面
		this.list.MoveToFront(node)
		return
	}

	// 不存在 把这本书放在最前面
	this.keyToNode[key] = this.list.PushFront(entry{key: key, value: value})
	if len(this.keyToNode) > this.cap {
		// 容量太多
		delete(this.keyToNode, this.list.Remove(this.list.Back()).(entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
