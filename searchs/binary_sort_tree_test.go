package searches
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestBinarySortTree(t *testing.T)  {
    items := []int{ 62,88,58,47,35,73,51,99,37,93}
    bst := NewBinarySortTree(items[0])
    for i := 1; i < len(items); i++ {
        bst.Insert(items[i])
    }
    Convey("#Insert", t, func() {
       fmt.Printf("二叉排序树: ")
       PreOrderPrint(bst.root)
       fmt.Printf("二叉排序树前序遍历: ")
       bst.root.PreOrderTraverse()
       fmt.Printf("\n二叉排序树中序遍历: ")
       bst.root.InOrderTraverse()
       fmt.Printf("\n二叉排序树后序遍历: ")
       bst.root.PostOrderTraverse()
       fmt.Printf("\n二叉排序树查找: ")
       node := bst.Search(37)
       fmt.Println(node.val)
    })

    Convey("#GetMaxNode", t, func() {
       fmt.Printf("二叉排序树查找 47 子结点的最大结点: ")
       node := bst.Search(47).GetMaxNode()
       fmt.Println(node.val)
    })

    Convey("#GetMinNode", t, func() {
       fmt.Printf("\n二叉排序树查找 47 子结点的最小结点: ")
       node := bst.Search(47).GetMinNode()
       fmt.Println(node.val)
    })

    Convey("#Delete", t, func() {
      fmt.Printf("\n二叉排序树中序遍历: ")
       bst.root.InOrderTraverse()
       bst.Delete(47)
       fmt.Printf("\n二叉排序树删除 47 后：\n")
       PreOrderPrint(bst.root)
       fmt.Printf("二叉排序树中序遍历: ")
       bst.root.InOrderTraverse()
    })
}