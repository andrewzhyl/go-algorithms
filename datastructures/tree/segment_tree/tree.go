package segment_tree

import (
	"fmt"
	"strings"
)

type Elem interface{}

type Merger func(a, b Elem) Elem

type SegmentTree struct {
	nodes []Elem
	data  []Elem
	size  int
	merge Merger
}

func NewSegmentTree(from []Elem, merge Merger) *SegmentTree {
	treeSize := len(from) * 4
	nodes := make([]Elem, treeSize)

	t := &SegmentTree{nodes, from, len(from), merge}
	t.Build(from, 0, 0, len(from)-1)
	return t
}

func (t *SegmentTree) GetSize() int {
	return len(t.nodes)
}

// 返回一个索引所表示的元素的左子树结点的索引
func (m *SegmentTree) leftchild(index int) int {
	return 2*index + 1
}

// 返回一个索引所表示的元素的右子树结点的索引
func (m *SegmentTree) rightchild(index int) int {
	return 2*index + 2
}

func (t *SegmentTree) Build(from []Elem, treeIndex, leftBound, rightBound int) {
	if leftBound == rightBound {
		t.nodes[treeIndex] = from[leftBound]
		return
	}
	leftTreeIndex := t.leftchild(treeIndex)
	rightTreeIndex := t.rightchild(treeIndex)
	mid := leftBound + (rightBound-leftBound)/2
	t.Build(from, leftTreeIndex, leftBound, mid)
	t.Build(from, rightTreeIndex, mid+1, rightBound)
	t.nodes[treeIndex] = t.merge(t.nodes[leftTreeIndex], t.nodes[rightTreeIndex])
}

func (t *SegmentTree) String() string {
	var ret = make([]string, len(t.nodes))
	for i := 0; i < len(t.nodes); i++ {
		if t.nodes[i] != nil {
			ret[i] = fmt.Sprintf("%d", t.nodes[i])
		} else {
			ret[i] = "nil"
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(ret, ","))
}

func (t *SegmentTree) Query(queryL, queryR int) Elem {
	if queryL < 0 ||
		queryL >= len(t.data) ||
		queryR < 0 ||
		queryR >= len(t.data) || queryL > queryR {
		panic("Index is illegal.")
	}
	return t.rangeQuery(0, 0, len(t.data)-1, queryL, queryR)
}

// 在以 treeID 为根的线段树中{leftBound...rightBound}的范围里，搜索区间[queryL..queryR]的值
func (t *SegmentTree) rangeQuery(treeIndex, leftBound, rightBound, queryL, queryR int) Elem {
	if leftBound == queryL && rightBound == queryR {
		return t.nodes[treeIndex]
	}
	mid := leftBound + (rightBound-leftBound)/2
	leftTreeIndex := t.leftchild(treeIndex)
	rightTreeIndex := t.rightchild(treeIndex)

	if queryL >= mid+1 {
		return t.rangeQuery(rightTreeIndex, mid+1, rightBound, queryL, queryR)
	} else if queryR <= mid {
		return t.rangeQuery(leftTreeIndex, leftBound, mid, queryL, queryR)
	}
	leftResult := t.rangeQuery(leftTreeIndex, leftBound, mid, queryL, mid)
	rightResult := t.rangeQuery(rightTreeIndex, mid+1, rightBound, mid+1, queryR)
	return t.merge(leftResult, rightResult)
}
func (t *SegmentTree) Set(index int, e Elem) {
	if index < 0 || index >= len(t.data) {
		panic("Index is illegal.")
	}
	t.update(0, 0, len(t.data)-1, index, e)
}
func (t *SegmentTree) update(treeIndex, leftBound, rightBound, index int, e Elem) {
	if leftBound == rightBound {
		t.nodes[treeIndex] = e
		return
	}
	mid := leftBound + (rightBound-leftBound)/2
	leftTreeIndex := t.leftchild(treeIndex)
	rightTreeIndex := t.rightchild(treeIndex)
	if index >= mid+1 {
		t.update(rightTreeIndex, mid+1, rightBound, index, e)
	} else { // index <= mid
		t.update(leftTreeIndex, leftBound, mid, index, e)
	}

	t.nodes[treeIndex] = t.merge(t.nodes[leftTreeIndex], t.nodes[rightTreeIndex])
}
