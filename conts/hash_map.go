package conts

import (
	"hash/maphash"
	"iter"

	"github.com/withoutforget/gods/clibutils"
)

type mapState = int32

var (
	stateActive  mapState = 0
	stateDeleted mapState = 1
	stateEmpty   mapState = 2
)

type hashNode[K comparable, V any] struct {
	key   K
	value V
	state mapState
}

func defaulthashNode[K comparable, V any]() hashNode[K, V] {
	return hashNode[K, V]{state: stateEmpty}
}

type HashMap[K comparable, V any] struct {
	data         List[hashNode[K, V]]
	elementCount int
	deletedCount int
	seed         maphash.Seed
}

func NewHashMap[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{seed: maphash.MakeSeed()}
}

func newHashMap[K comparable, V any](seed maphash.Seed, cap int) HashMap[K, V] {
	var h = HashMap[K, V]{
		data: NewList[hashNode[K, V]](cap),
		seed: seed,
	}
	for range cap {
		h.data.Append(defaulthashNode[K, V]())
	}
	return h
}

func (this *HashMap[K, V]) LoadFactor() float32 {
	if this.data.Len() == 0 {
		return 1
	}
	return float32(this.elementCount) / float32(this.data.Len())
}

func (this *HashMap[K, V]) TombStoneFactor() float32 {
	return float32(this.deletedCount) / float32(this.data.Len())
}

func (this *HashMap[K, V]) hash(key K) int {
	var h = clibutils.HashFunction(this.seed, key)
	if h < 0 {
		h = -h
	}
	return h % this.data.Len()
}

func (this *HashMap[K, V]) swap(rhs *HashMap[K, V]) {
	this.data, rhs.data = clibutils.Swap(this.data, rhs.data)
	this.elementCount, rhs.elementCount = clibutils.Swap(this.elementCount, rhs.elementCount)
	this.deletedCount, rhs.deletedCount = clibutils.Swap(this.deletedCount, rhs.deletedCount)
}

func (this *HashMap[K, V]) realloc(newCap int) {
	var tmp = newHashMap[K, V](this.seed, newCap)
	for i := range this.data.All() {
		if i.state == stateActive {
			tmp.Set(i.key, i.value)
		}
	}
	this.swap(&tmp)
}

func (this *HashMap[K, V]) Set(key K, value V) {
	if this.LoadFactor() > 0.75 {
		if this.elementCount != 0 {
			this.realloc(this.elementCount * 2)
		} else {
			this.realloc(1)
		}
	}

	var startPos = this.hash(key)
	var firstTombstone = -1

	for i := range this.data.Len() {
		var pos = (startPos + i) % this.data.Len()
		node := this.data.Get(pos)

		switch node.state {
		case stateActive:
			if node.key == key {
				node.value = value
				return
			}
		case stateDeleted:
			if firstTombstone == -1 {
				firstTombstone = pos
			}
		case stateEmpty:
			if firstTombstone != -1 {
				pos = firstTombstone
				node = this.data.Get(pos)
				this.deletedCount--
			}
			node.key = key
			node.value = value
			node.state = stateActive
			this.elementCount++
			return
		}
	}

	node := this.data.Get(firstTombstone)
	node.key = key
	node.value = value
	node.state = stateActive
	this.elementCount++
	this.deletedCount--
}

func (this *HashMap[K, V]) Get(key K) *V {
	if this.elementCount == 0 {
		return nil
	}
	var startPos = this.hash(key)
	for i := range this.data.Len() {
		var pos = (startPos + i) % this.data.Len()
		node := this.data.Get(pos)
		if node.state == stateEmpty {
			return nil
		}
		if node.state == stateActive {
			if node.key == key {
				return &node.value
			}
		}
	}
	return nil
}

func (this *HashMap[K, V]) Delete(key K) {
	if this.elementCount == 0 {
		return
	}
	var startPos = this.hash(key)
	for i := range this.data.Len() {
		var pos = (startPos + i) % this.data.Len()
		node := this.data.Get(pos)
		if node.state == stateActive {
			if node.key == key {
				node.state = stateDeleted
				this.deletedCount++
				this.elementCount--
				break
			}
		}
		if node.state == stateEmpty {
			break
		}
	}
	if this.TombStoneFactor() > 0.3 {
		this.realloc(this.data.Len())
	}
}

func (this *HashMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for n := range this.data.All() {
			if n.state != stateActive {
				continue
			}
			if !yield(n.key) {
				return
			}
		}
	}
}

func (this *HashMap[K, V]) Values() iter.Seq[*V] {
	return func(yield func(*V) bool) {
		for n := range this.data.All() {
			if n.state != stateActive {
				continue
			}
			if !yield(&n.value) {
				return
			}
		}
	}
}

func (this *HashMap[K, V]) All() iter.Seq2[K, *V] {
	return func(yield func(K, *V) bool) {
		for n := range this.data.All() {
			if n.state == stateActive {
				if !yield(n.key, &n.value) {
					return
				}
			}
		}
	}
}

func (this *HashMap[K, V]) Len() int {
	return this.elementCount
}
