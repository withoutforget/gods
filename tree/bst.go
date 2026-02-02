package tree

import "iter"

type BstNode[T comparable] struct {
	value T
	left  *BstNode[T]
	right *BstNode[T]
}

type BST[T comparable] struct {
	head *BstNode[T]
	comp func(a, b T) bool // less
}

func NewBST[T comparable](comp func(a, b T) bool) BST[T] {
	return BST[T]{comp: comp}
}

func (this *BST[T]) Insert(value T) {
	if this.head == nil {
		this.head = &BstNode[T]{value: value}
		return
	}
	parent := this.head
	child := parent
	for child != nil {
		if this.comp(child.value, value) {
			parent = child
			child = child.right
		} else {
			parent = child
			child = child.left
		}
	}
	var node = &BstNode[T]{value: value}
	if this.comp(parent.value, value) {
		parent.right = node
	} else {
		parent.left = node
	}
}

func (this *BST[T]) searchNode(value T) (Parent *BstNode[T], Child *BstNode[T]) {
	if this.head == nil {
		return nil, nil
	}
	parent := this.head
	child := parent
	for child != nil {
		if !this.comp(child.value, value) && !this.comp(value, child.value) {
			if parent == child {
				return nil, child
			}
			return parent, child
		}
		if this.comp(child.value, value) {
			parent = child
			child = child.right
		} else {
			parent = child
			child = child.left
		}
	}
	return nil, nil
}

func (this *BST[T]) Search(value T) bool {
	var a, b = this.searchNode(value)
	return a != nil && b != nil
}

func (this *BST[T]) Delete(value T) {
	parent, node := this.searchNode(value)
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		if parent == nil {
			this.head = nil
			return
		}
		if parent.left == node {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}
	if node.left == nil || node.right == nil {
		var replaceNode *BstNode[T]
		if node.left != nil {
			replaceNode = node.left
		} else {
			replaceNode = node.right
		}
		if parent == nil {
			this.head = replaceNode
			return
		}
		if parent.left == node {
			parent.left = replaceNode
		} else {
			parent.right = replaceNode
		}
		return
	}

	minParent := node
	minNode := node.right
	for minNode.left != nil {
		minParent = minNode
		minNode = minNode.left
	}
	node.value = minNode.value
	if minParent.left == minNode {
		minParent.left = minNode.right
	} else {
		minParent.right = minNode.right
	}
}

func (this *BST[T]) dfs(node *BstNode[T], arr *[]T) {
	if node == nil {
		return
	}
	this.dfs(node.left, arr)
	*arr = append(*arr, node.value)
	this.dfs(node.right, arr)
}

func (this *BST[T]) All() iter.Seq[T] {
	var output = make([]T, 0)
	this.dfs(this.head, &output)
	return func(yield func(T) bool) {
		for _, v := range output {
			if !yield(v) {
				return
			}
		}
	}
}
