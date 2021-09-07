package bstmap

import (
	"fmt"
	"strings"
)

type K interface{}
type V interface{}

type Node struct {
	key   K
	value V
	left  *Node
	right *Node
}

func NewNode(key K, value V) *Node {
	return &Node{key, value, nil, nil}
}

// 二叉查找树
type BSTMap struct {
	Root *Node
	size int
}

func NewBSTMap() *BSTMap {
	return &BSTMap{nil, 0}
}

// 获取大小
func (bst *BSTMap) GetSize() int {
	return bst.size
}

// 是否为空
func (bst *BSTMap) IsEmpty() bool {
	return bst.size == 0
}

// 插入操作
func (bst *BSTMap) Add(key, val interface{}) *BSTMap {
	if bst.Root == nil {
		bst.Root = NewNode(key, val)
	} else {
		bst.Root.add(key, val)
	}
	bst.size++
	return bst
}

// 递归插入
func (n *Node) add(key, val interface{}) *Node {
	if n == nil {
		return NewNode(key, val)
	}
	if compare(key, n.key) < 0 {
		n.left = n.left.add(key, val)
	} else if compare(key, n.key) > 0 {
		n.right = n.right.add(key, val)
	} else {
		n.value = val
	}
	return n
}

// 二叉树查找
func (t *BSTMap) GetNode(key interface{}) *Node {
	if t == nil || t.Root == nil {
		return nil
	}
	return t.Root.getNode(key)
}

// 查找
func (n *Node) getNode(key interface{}) *Node {
	if n == nil {
		return nil
	}

	if compare(key, n.key) == 0 {
		return n
	}
	if compare(key, n.key) < 0 {
		return n.left.getNode(key)
	}
	return n.right.getNode(key)
}

func (t *BSTMap) Contains(key interface{}) bool {
	return t.GetNode(key) != nil
}

func (ll *BSTMap) Get(key interface{}) interface{} {
	node := ll.GetNode(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (ll *BSTMap) Set(key, val interface{}) {
	node := ll.GetNode(key)
	if node == nil {
		panic(key.(string) + " doesn't exist!")
	}
	node.value = val
}

// 二叉树删除
func (t *BSTMap) Delete(key interface{}) *Node {
	// 如果为空，直接返回
	if t == nil || t.Root == nil {
		return nil
	}

	ret := t.GetNode(key)
	if ret == nil {
		return nil
	}
	t.size--
	return t.Root.Delete(key)
}

// 二叉树删除
func (node *Node) Delete(key interface{}) *Node {
	if node == nil {
		return node
	}

	// 如果删除的值小于 node 的值，则递归的遍历 node 的左子树
	// 删除的结点位于左子树
	if compare(key, node.key) < 0 {
		node.left = node.left.Delete(key)

		// 如果删除的值大于 node 的值，则递归的遍历 node 的右子树
		// 删除的结点位于右子树
	} else if compare(key, node.key) > 0 {
		node.right = node.right.Delete(key)

		// 如果值相等，则删除结点
	} else {
		// 一、被删除的结点是叶子结点：直接从二叉排序树当中移除即可
		// 二、被删除的结点有 1 个子结点，子结点需要连接到父亲结点
		if node.left == nil {
			node = node.right
			return node
		} else if node.right == nil {
			node = node.left
			return node
		}

		// 三、被删除的结点有左右子结点都不为空

		// 获取删除结点中序遍历的直接前驱结点
		temp := node.left.GetMaxNode()

		// 将删除结点的值替换为直接后继结点的值
		node.key = temp.key
		node.value = temp.value

		// 删除直接后继结点
		node.left = node.left.Delete(temp.key)

		// 或者，获取删除结点中序遍历的直接后驱结点
		// temp := node.right.GetMinNode()
		// node.key = temp.key
		// node.value = temp.value
		// node.right = node.right.Delete(temp.key)

	}
	return node
}

// 获取最大结点
func (node *Node) GetMaxNode() *Node {
	if node == nil {
		return nil
	}
	for {
		if node.right == nil {
			return node
		}
		node = node.right
	}
}

// 获取最小结点
func (node *Node) GetMinNode() *Node {
	if node == nil {
		return nil
	}
	for {
		if node.left == nil {
			return node
		}
		node = node.left
	}
}

// 前序遍历打印
func (n *Node) PreOrderTraverse() {
	if n == nil {
		return
	}
	fmt.Printf(" -> %v:%d \n", n.key, n.value)
	if n.left != nil {
		n.left.PreOrderTraverse()
	}
	if n.right != nil {
		n.right.PreOrderTraverse()
	}
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
