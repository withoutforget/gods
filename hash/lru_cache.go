package hash

import (
	"iter"

	"github.com/withoutforget/gods/list"
)

type nodeState int

var (
	nodeEmpty  nodeState = 0
	nodeActive nodeState = 1
)

type lruNode[K comparable, T any] struct {
	key   K
	value T
	state nodeState
}

type LRU[K comparable, V any] struct {
	linkedList list.LinkedList[lruNode[K, V]]
	hashMap    HashMap[K, *list.LinkedListNode[lruNode[K, V]]]
	cap        int
}

func NewLRU[K comparable, V any](cap int) *LRU[K, V] {
	var lru LRU[K, V]
	for range cap {
		lru.linkedList.PushBack(lruNode[K, V]{state: nodeState(nodeEmpty)})
	}
	lru.cap = cap

	return &lru
}

func (this *LRU[K, V]) Put(key K, value V) {
	var node *list.LinkedListNode[lruNode[K, V]]
	n := this.hashMap.Get(key)
	if n == nil {
		node = this.linkedList.Back()
		nval := node.GetValue()
		if nval.state == nodeActive {
			this.hashMap.Delete(nval.key)
		}
	} else {
		node = *n
	}
	this.linkedList.Erase(node)
	this.linkedList.PushFront(lruNode[K, V]{key: key, value: value, state: nodeState(nodeActive)})
	this.hashMap.Set(key, this.linkedList.Front())
}

func (this *LRU[K, V]) Get(key K) *V {
	var node = this.hashMap.Get(key)
	if node == nil {
		return nil
	}
	nval := *node
	this.linkedList.Erase(*node)
	this.linkedList.PushFront(lruNode[K, V]{key: key, value: nval.GetValue().value, state: nodeState(nodeActive)})
	this.hashMap.Set(key, this.linkedList.Front())
	return &nval.GetValue().value
}

func (this *LRU[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range this.linkedList.All() {
			if v.state != nodeActive {
				continue
			}
			if !yield(v.key, v.value) {
				return
			}
		}
	}
}
