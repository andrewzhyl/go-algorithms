package binarytree

import (
	"fmt"
	"testing"

	// "reflect"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsert(t *testing.T) {
	Convey("#Insert", t, func() {
		tree := NewTree()
		tree.Insert(100).
			Insert(-20).
			Insert(-50).
			Insert(-15).
			Insert(-60).
			Insert(50).
			Insert(60).
			Insert(55).
			Insert(85).
			Insert(15).
			Insert(5).
			Insert(-10)
		fmt.Print("\n")
		PreOrder(tree.Root, 0, 'M')
		tree.Root.PreOrderTraverse()
	})
}

func TestTree(t *testing.T) {
	tree := NewTree()
	root := &BinaryTreeNode{"A", nil, nil}
	nb := NewBinaryTreeNode("B")
	nc := NewBinaryTreeNode("C")
	nd := NewBinaryTreeNode("D")
	ne := NewBinaryTreeNode("E")
	nf := NewBinaryTreeNode("F")
	ng := NewBinaryTreeNode("G")
	nh := NewBinaryTreeNode("H")
	ni := NewBinaryTreeNode("I")

	tree.Root = root
	root.left = nb
	root.right = nc
	root.left.left = nd
	root.left.right = ne
	root.right.left = nf
	root.right.right = ng
	root.left.left.left = nh
	root.left.left.right = ni
	fmt.Print("\n")
	PreOrder(root, 0, 'M')

	Convey("#PreOrderTraverse", t, func() {
		fmt.Printf("Pre Order Traversal of the given tree is: ")
		root.PreOrderTraverse()
	})

	Convey("#InOrderTraverse", t, func() {
		fmt.Printf("In Order Traversal of the given tree is: ")
		root.InOrderTraverse()
	})

	Convey("#PostOrderTraverse", t, func() {
		fmt.Printf("Post Order Traversal of the given tree is: ")
		root.PostOrderTraverse()
	})

	Convey("#PreOrderSearch", t, func() {
		fmt.Printf("前序遍历搜索: ")
		fmt.Println(tree.PreOrderSearch(tree.Root, "G"))
	})

	Convey("#InOrderSearch", t, func() {
		fmt.Printf("中序遍历搜索: ")
		fmt.Println(root.InOrderSearch("C"))
	})

	Convey("#PostOrderSearch", t, func() {
		fmt.Printf("后续遍历搜索: ")
		fmt.Println(root.PostOrderSearch("D"))
	})
}

func TestI2num(t *testing.T) {
	Convey("#I2num", t, func() {
		So(I2num('A'), ShouldEqual, 65)
		So(I2num("A"), ShouldEqual, 65)
		So(I2num(100), ShouldEqual, 100)
	})
}
