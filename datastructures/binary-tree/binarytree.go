package binarytree

import(
	"fmt"
	"reflect"
)

type Node struct{
	val interface{}
	left *Node
	right *Node
}

type Tree struct{
	Root *Node
}

func NewNode(val interface{}) *Node {
	return &Node{ val, nil, nil }
}

func NewTree() *Tree {
	return &Tree{ nil }
}

func (t *Tree) Insert(val interface{}) *Tree{
	if t.Root == nil {
		t.Root = NewNode(val)
	} else{
		t.Root.Insert(val)
	}
	return t
}

func (n *Node) Insert(val interface{}){
	if(n == nil){
		return 
	} 
	var terminalNode *Node
	for n != nil {
		terminalNode = n
		if int(val.(int)) <= int(n.val.(int)) {
			n = n.left
		} else {
			n = n.right
		}
	}

	if int(val.(int)) <= int(terminalNode.val.(int)) {
		terminalNode.left = NewNode(val)
	} else {
		terminalNode.right = NewNode(val)
	}
	return
}

// 前序遍历递归
func (n *Node) PreOrderTraverse() {
	fmt.Print(" -> ", n.val)
	if n.left != nil{
		n.left.PreOrderTraverse()
	}
	if n.right != nil{
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

	if n.left != nil{
		 v1 := b.PreOrderSearch(n.left, val)
		 if v1 != nil {
		 	return v1
		 }
	}
	if n.right != nil{
		 v2 := b.PreOrderSearch(n.right, val)
		 if v2 != nil {
		 	return v2
		 }
	}
	return nil
}

// 前序遍历打印
func PreOrder(node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v\n", ch, node.val)
	PreOrder(node.left, ns+2, 'L')
	PreOrder(node.right, ns+2, 'R')
}

// 中序遍历递归
func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}
	if n.left != nil{
		n.left.InOrderTraverse()
	}
	fmt.Print(" -> ", n.val)
	if n.right != nil{
		n.right.InOrderTraverse()
	}
}

// 中序遍历查找
func (n *Node) InOrderSearch(val interface{}) interface{} {
	if n.left != nil{
		l1 := n.left.InOrderSearch(val)
		if l1 != nil {
			return l1
		}
	}
	if n.val == val {
		return n.val
	}
	if n.right != nil{
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
	if n.left != nil{
		n.left.PostOrderTraverse()
	}
	if n.right != nil{
		n.right.PostOrderTraverse()
	}
	fmt.Print(" -> ", n.val)
}

// 后序遍历查找
func (n *Node) PostOrderSearch(val interface{}) interface{} {
	if n.left != nil{
		l1 := n.left.PostOrderSearch(val)
		if l1 != nil {
			return l1
		}
	}
	if n.right != nil{
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

func I2num(val interface{}) int { // interface to number
	var i int
	switch val.(type) {
    case string:
    	i = int(reflect.ValueOf(val).String()[0])
    case int32:
    	i  =  int(val.(int32))
    default:
        i  =  int(val.(int))
    }
    return i
}