package avltree

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	opAdd = iota
	opRemove
	opSearch
)

func TestAVLTree(t *testing.T) {

	var tree = NewTree()
	fmt.Println("Empty Tree:")
	avl, _ := json.MarshalIndent(tree, "", "   ")
	fmt.Println(string(avl))
	Convey("Insert Tree", t, func() {
		fmt.Println("\nInsert Tree:")
		tree.Add(4, 4)
		tree.Add(2, 2)
		tree.Add(7, 7)
		tree.Add(6, 6)
		tree.Add(6, 6)
		tree.Add(9, 9)
		avl, _ = json.MarshalIndent(tree, "", "   ")
		tree.Root.InOrderTraverse()
	})

	Convey("Remove Tree", t, func() {
		fmt.Println("\nRemove Tree:")
		tree.Delete(4)
		tree.Delete(6)
		avl, _ = json.MarshalIndent(tree, "", "   ")
		tree.Root.InOrderTraverse()
	})
}
