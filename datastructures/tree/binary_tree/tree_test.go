package binarytree

import (
	"fmt"
	"testing"

	// "reflect"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsert(t *testing.T) {
	Convey("#Insert", t, func() {

		items := []int{5, 6, 2, 4, 1, 8, 7, 9, 3}
		tree := NewTree()
		for i := 0; i < len(items); i++ {
			tree.Insert(items[i])
		}
		fmt.Printf("二叉树中序遍历: ")
		tree.Root.InOrderTraverse()
		fmt.Printf("\n")
		PreOrder(tree.Root, 0, 'M')
	})
}

func TestTree(t *testing.T) {
	items := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	tree := NewTree()
	for i := 0; i < len(items); i++ {
		tree.Insert(items[i])
	}
	root := tree.Root
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
		fmt.Printf("%c", tree.PreOrderSearch(tree.Root, 'G'))
	})

	Convey("#InOrderSearch", t, func() {
		fmt.Printf("中序遍历搜索: ")
		fmt.Printf("%c", root.InOrderSearch('C'))
	})

	Convey("#PostOrderSearch", t, func() {
		fmt.Printf("后续遍历搜索: ")
		fmt.Printf("%c", root.PostOrderSearch('D'))
	})
}
