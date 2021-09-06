package maxheap

import (
	"strings"
)

type Elem interface{}

type Maxheap struct {
	heapArray []Elem
}

func NewMaxHeap() *Maxheap {
	maxheap := &Maxheap{
		heapArray: []Elem{},
	}
	return maxheap
}

func (m *Maxheap) size() int {
	return len(m.heapArray)
}

func (m *Maxheap) isEmpty() bool {
	return len(m.heapArray) == 0
}

// 返回一个索引所表示的元素的父亲结点的索引
func (m *Maxheap) parent(index int) int {
	if index == 0 {
		panic("index-0 doesn't have parent. ")
	}
	return (index - 1) / 2
}

// 返回一个索引所表示的元素的左子树结点的索引
func (m *Maxheap) leftchild(index int) int {
	return 2*index + 1
}

// 返回一个索引所表示的元素的右子树结点的索引
func (m *Maxheap) rightchild(index int) int {
	return 2*index + 2
}

func (m *Maxheap) isLeaf(index int) bool {
	if m.leftchild(index) > m.size() {
		return true
	}
	return false
}

func (m *Maxheap) add(item Elem) error {
	// if m.size() >= m.maxsize {
	// 	return errors.New("Heal is ful")
	// }
	// 增加到最后
	m.heapArray = append(m.heapArray, item)
	m.siftup(m.size() - 1)
	return nil
}

// 上浮元素
func (m *Maxheap) siftup(index int) {
	for index > 0 && compare(m.heapArray[m.parent(index)], m.heapArray[index]) < 0 {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

// 交换位置
func (m *Maxheap) swap(a, b int) {
	m.heapArray[a], m.heapArray[b] = m.heapArray[b], m.heapArray[a]
}

// 下沉元素
func (m *Maxheap) siftdown(index int) {
	if m.isLeaf(index) {
		return
	}

	largest := index
	leftChildIndex := m.leftchild(index)
	rightChildIndex := m.rightchild(index)

	//If index is smallest then return
	if leftChildIndex < m.size() && compare(m.heapArray[leftChildIndex], m.heapArray[largest]) > 0 {
		largest = leftChildIndex
	}

	if rightChildIndex < m.size() && compare(m.heapArray[rightChildIndex], m.heapArray[largest]) > 0 {
		largest = rightChildIndex
	}
	if largest != index {
		m.swap(largest, index)
		m.siftdown(largest)
	}
}

// 取出堆中最大元素
func (m *Maxheap) ExceptMax() Elem {
	if m.size() == 0 {
		panic("Can not findMax when heap is empty!")
	}
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size()-1]
	m.heapArray = m.heapArray[:m.size()-1]
	m.siftdown(0)
	return top
}

func compare(a, b interface{}) int {
	switch a.(type) {
	case string:
		return strings.Compare(a.(string), b.(string))
	case int32:
		return func(a, b int32) int {
			if a == b {
				return 0
			}
			if a < b {
				return -1
			}
			return +1
		}(a.(int32), b.(int32))
	default:
		return func(a, b int) int {
			if a == b {
				return 0
			}
			if a < b {
				return -1
			}
			return +1
		}(a.(int), b.(int))
	}
}
