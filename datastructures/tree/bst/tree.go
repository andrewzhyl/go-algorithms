package bst

import (
	"fmt"
	"strings"
)

type Node struct {
	val   interface{}
	left  *Node
	right *Node
}

// 二叉查找树
type BinarySortTree struct {
	Root *Node
	size int
}

func NewBinarySortTree() *BinarySortTree {
	return &BinarySortTree{nil, 1}
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}

// 插入操作
func (bst *BinarySortTree) Insert(val interface{}) *BinarySortTree {
	ret := bst.Search(val)
	if ret != nil {
		return bst
	}
	if bst.Root == nil {
		bst.Root = NewNode(val)
	} else {
		bst.Root.Insert(val)
	}
	bst.size++
	return bst
}

// 递归插入
func (n *Node) Insert(val interface{}) *Node {
	if n == nil {
		return NewNode(val)
	}
	if compare(val, n.val) < 0 {
		n.left = n.left.Insert(val)
	} else if compare(val, n.val) > 0 {
		n.right = n.right.Insert(val)
	}
	return n
}

// IsEmpty checks if the stack is empty
func (s *BinarySortTree) IsEmpty() bool {
	return s.size == 0
}

// 二叉树查找
func (t *BinarySortTree) Search(val interface{}) *Node {
	if t == nil || t.Root == nil {
		return nil
	}
	return t.Root.Search(val)
}

// 查找
func (n *Node) Search(key interface{}) *Node {
	if n == nil {
		return nil
	}

	if compare(key, n.val) == 0 {
		return n
	}
	if compare(key, n.val) < 0 {
		return n.left.Search(key)
	}
	return n.right.Search(key)
}

// 二叉树删除
func (t *BinarySortTree) Delete(key interface{}) {
	// 如果为空，直接返回
	if t == nil || t.Root == nil {
		return
	}

	ret := t.Search(key)
	if ret == nil {
		return
	}

	t.size--
	t.Root.Delete(key)
}

// 二叉树删除
func (node *Node) Delete(key interface{}) *Node {
	if node == nil {
		return node
	}

	// 如果删除的值小于 node 的值，则递归的遍历 node 的左子树
	// 删除的结点位于左子树
	if compare(key, node.val) < 0 {
		node.left = node.left.Delete(key)

		// 如果删除的值大于 node 的值，则递归的遍历 node 的右子树
		// 删除的结点位于右子树
	} else if compare(key, node.val) > 0 {
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
		node.val = temp.val

		// 删除直接后继结点
		node.left = node.left.Delete(temp.val)

		// 或者，获取删除结点中序遍历的直接后区结点
		// temp := node.right.GetMinNode()
		// node.val = temp.val
		// node.right = node.right.Delete(temp.val)

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

// 获取最小结点
func (bst *BinarySortTree) GetSize() int {
	return bst.size
}

// 前序遍历打印
func (n *Node) PreOrderTraverse() {
	if n == nil {
		return
	}
	fmt.Print(" -> ", n.val)
	if n.left != nil {
		n.left.PreOrderTraverse()
	}
	if n.right != nil {
		n.right.PreOrderTraverse()
	}
}

// 中序遍历打印
func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.InOrderTraverse()
	}
	fmt.Print(" -> ", n.val)
	if n.right != nil {
		n.right.InOrderTraverse()
	}
}

// 后序遍历打印
func (n *Node) PostOrderTraverse() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.PostOrderTraverse()
	}
	if n.right != nil {
		n.right.PostOrderTraverse()
	}
	fmt.Print(" -> ", n.val)
}

// Print the Huffman tree
func PreOrderPrint(Root *Node) {
	fmt.Printf("\n")
	var traverse func(node *Node, ns int, ch rune)
	traverse = func(node *Node, ns int, ch rune) {
		if node == nil {
			return
		}
		for i := 0; i < ns; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%c:%v\n", ch, node.val)
		traverse(node.left, ns+2, 'L')
		traverse(node.right, ns+2, 'R')
	}
	traverse(Root, 0, 'M')
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
