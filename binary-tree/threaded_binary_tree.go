package binarytree

import(
	"fmt"
	"reflect"
)

type Node struct{
	val interface{}
	left *Node
	right *Node
	prev  *Node
	next  *Node
}

type Tree struct{
	Root *Node
}

// 全局变量，始终指向刚刚访问过的结点
var thread *Node 

func NewNode(val interface{}) *Node {
	return &Node{ val, nil, nil, nil, nil }
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

func (n *Node) Insert(val interface{}) *Node{
	if(n == nil){
		return nil
	} else if I2num(val) <= I2num(n.val) {
		if n.left == nil {
			n.left = NewNode(val)
		}else{
			n.left.Insert(val)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(val)
		}else{
			n.right.Insert(val)
		}
	}
	return n
}

// 前序遍历递归
func (n *Node) PreOrderTraverse() {
	fmt.Printf("%v", n.val)
	if n.left != nil{
		n.left.PreOrderTraverse()
	}
	if n.right != nil{
		n.right.PreOrderTraverse()
	}
}

//中序遍历，中序线索化二叉树
func (t *Tree) threaded() {
	thread = nil
	t.InOrderClue(t.Root)
}

//中序遍历，中序线索化二叉树，将二叉树变成一个双向链表
func (t *Tree) InOrderClue(n *Node) {
	if n == nil {
		return
	}
	t.InOrderClue(n.left)
	if n.prev == nil{
		n.prev = thread
	}
	if thread != nil && thread.next == nil{
		thread.next = n
	}
	thread = n
	t.InOrderClue(n.right)
}

func (t *Tree) printThreadedTree(){
	fmt.Println("线索化二叉树：")
	node := t.Root
	for node.left != nil {
		node = node.left
	}
	for( node != nil){
		fmt.Print(" -> ", node.val)
		node = node.next
	}
	fmt.Print("\n")
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
