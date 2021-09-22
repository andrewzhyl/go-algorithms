package binarytree

import (
	"fmt"
	"strings"
)

type Node struct {
	val   interface{}
	left  *Node
	right *Node
}

type Tree struct {
	Root *Node
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}

func NewTree() *Tree {
	return &Tree{nil}
}

func (t *Tree) Insert(val interface{}) *Tree {
	if t.Root == nil {
		t.Root = NewNode(val)
	} else {
		t.Root.Insert(val)
	}
	return t
}

func (node *Node) Insert(val interface{}) {
	if node == nil {
		return
	}
	var cur *Node
	for node != nil {
		cur = node
		if compare(val, node.val) <= 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	if compare(val, cur.val) <= 0 {
		cur.left = NewNode(val)
	} else {
		cur.right = NewNode(val)
	}
	return
}

// 前序遍历递归
func (n *Node) PreOrderTraverse() {
	fmt.Printf(" -> %s", convertToString(n.val))
	if n.left != nil {
		n.left.PreOrderTraverse()
	}
	if n.right != nil {
		n.right.PreOrderTraverse()
	}
}

// 前序遍历查找
func (b *Tree) PreOrderSearch(n *Node, val interface{}) interface{} {
	if n == nil {
		return nil
	}
	if n.val == val {
		return n.val
	}

	if n.left != nil {
		v1 := b.PreOrderSearch(n.left, val)
		if v1 != nil {
			return v1
		}
	}
	if n.right != nil {
		v2 := b.PreOrderSearch(n.right, val)
		if v2 != nil {
			return v2
		}
	}
	return nil
}

// 前序遍历打印
func PreOrder(n *Node, ns int, ch rune) {
	if n == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%s\n", ch, convertToString(n.val))
	PreOrder(n.left, ns+2, 'L')
	PreOrder(n.right, ns+2, 'R')
}

// 中序遍历递归
func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.InOrderTraverse()
	}
	fmt.Printf(" -> %s", convertToString(n.val))
	if n.right != nil {
		n.right.InOrderTraverse()
	}
}

// 中序遍历查找
func (n *Node) InOrderSearch(val interface{}) interface{} {
	if n.left != nil {
		l1 := n.left.InOrderSearch(val)
		if l1 != nil {
			return l1
		}
	}
	if n.val == val {
		return n.val
	}
	if n.right != nil {
		l2 := n.right.InOrderSearch(val)
		if l2 != nil {
			return l2
		}
	}
	return nil
}

// 后续遍历递归
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
	fmt.Printf(" -> %s", convertToString(n.val))
}

// 后序遍历查找
func (n *Node) PostOrderSearch(val interface{}) interface{} {
	if n.left != nil {
		l1 := n.left.PostOrderSearch(val)
		if l1 != nil {
			return l1
		}
	}
	if n.right != nil {
		l2 := n.right.PostOrderSearch(val)
		if l2 != nil {
			return l2
		}
	}

	if n.val == val {
		return n.val
	}
	return nil
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

func convertToString(a interface{}) string {
	switch a.(type) {
	case string:
		return a.(string)
	case rune:
		return fmt.Sprintf("%c", a)
	default:
		return fmt.Sprintf("%d", a)
	}
}
