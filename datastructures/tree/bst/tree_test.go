package bst

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySortTree(t *testing.T) {
	items := []int{5, 6, 2, 4, 1, 8, 7, 9, 3}
	bst := NewBinarySortTree()
	for i := 0; i < len(items); i++ {
		bst.Insert(items[i])
	}
	Convey("#Insert", t, func() {
		fmt.Printf("二叉排序树: ")
		PreOrderPrint(bst.Root)
		fmt.Printf("二叉排序树前序遍历: ")
		bst.Root.PreOrderTraverse()
		fmt.Printf("\n二叉排序树中序遍历: ")
		bst.Root.InOrderTraverse()
		fmt.Printf("\n二叉排序树后序遍历: ")
		bst.Root.PostOrderTraverse()
		fmt.Printf("\n二叉排序树查找: ")
		node := bst.Search(3)
		fmt.Println(node.val)
	})

	Convey("#GetMaxNode", t, func() {
		fmt.Printf("二叉排序树查找 8 子结点的最大结点: ")
		node := bst.Search(8).GetMaxNode()
		fmt.Println(node.val)
	})

	Convey("#GetMinNode", t, func() {
		fmt.Printf("\n二叉排序树查找 7 子结点的最小结点: ")
		node := bst.Search(8).GetMinNode()
		fmt.Println(node.val)
	})

	Convey("#Delete", t, func() {
		fmt.Printf("\n二叉排序树中序遍历: ")
		bst.Root.InOrderTraverse()
		bst.Delete(7)
		fmt.Printf("\n二叉排序树删除 7 后：\n")
		PreOrderPrint(bst.Root)
		fmt.Printf("二叉排序树中序遍历: ")
		bst.Root.InOrderTraverse()
	})
}
