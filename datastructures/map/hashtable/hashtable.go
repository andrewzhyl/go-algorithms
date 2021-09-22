package hashtable

import (
	"bytes"
	"fmt"
	"hash/fnv"
)

const MAP_SIZE = 50

type Node struct {
	key   string
	value interface{}
	next  *Node
}

type HashTable struct {
	Data []*Node
	size int
}

func NewHashTable() *HashTable {
	return &HashTable{Data: make([]*Node, MAP_SIZE)}
}

func hashFunction(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32()) % MAP_SIZE
}

func (h *HashTable) GetSize() int {
	return h.size
}

func (h *HashTable) Add(key string, val interface{}) {
	index := hashFunction(key)
	if h.Data[index] == nil {
		h.Data[index] = &Node{key: key, value: val}
		h.size++
	} else {
		head_node := h.Data[index]
		for ; head_node.next != nil; head_node = head_node.next {
			if head_node.key == key {
				head_node.value = val
				return
			}
		}
		head_node.next = &Node{key: key, value: val}
		h.size++
	}
}

func (h *HashTable) Set(key string, val interface{}) {
	index := hashFunction(key)
	if h.Data[index] == nil {
		panic(key + " doesn't exist!")
	}
	head_node := h.Data[index]
	for ; head_node != nil; head_node = head_node.next {
		if head_node.key == key {
			head_node.value = val
			return
		}
	}
	panic(key + " doesn't exist!")
}

func (h *HashTable) Get(key string) (interface{}, bool) {
	index := hashFunction(key)
	if h.Data[index] != nil {
		head_node := h.Data[index]
		for ; ; head_node = head_node.next {
			if head_node.key == key {
				return head_node.value, true
			}
			if head_node.next == nil {
				break
			}
		}
	}
	return nil, false
}

func (h *HashTable) Contains(key string) bool {
	_, has := h.Get(key)
	return has
}

func (h *HashTable) Remove(key string) bool {
	index := hashFunction(key)
	iterator := h.Data[index]
	if iterator == nil {
		return false
	}
	if iterator.key == key {
		h.Data[index] = iterator.next
		h.size--
		return true
	} else {
		prev := iterator
		iterator = iterator.next
		for iterator != nil {
			if iterator.key == key {
				prev.next = iterator.next
				h.size--
				return true
			}
			prev = iterator
			iterator = iterator.next
		}
	}
	return false
}

func (h *HashTable) String() string {
	var result bytes.Buffer
	fmt.Fprintln(&result, "{")
	for _, v := range h.Data {
		if v != nil {
			fmt.Fprintf(&result, "\t%s:%v\t\n", v.key, v.value)
			node := v.next
			for node != nil {
				fmt.Fprintf(&result, "\t%s:%v\t\n", node.key, node.value)
				node = node.next
			}
		}
	}
	fmt.Fprintln(&result, "}")
	return result.String()
}
