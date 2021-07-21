package binarytree
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestInsert(t *testing.T)  {
    Convey("#Insert", t, func() {
        tree := NewTree()
        tree.Insert(100)
        tree.Insert(-20)
        tree.Insert(-50)
        tree.Insert(-15)
        tree.Insert(-60)
        tree.Insert(50)
        tree.Insert(60)
        tree.Insert(55)
        tree.Insert(85)
        tree.Insert(15)
        tree.Insert(5)
        tree.Insert(-10)

        fmt.Print("\n")
        PreOrder(tree.Root, 0, 'M')
        tree.Root.PreOrderTraverse()
    })
}

func TestTree(t *testing.T)  {
    tree := NewTree()
    root :=NewNode("A")
    nb := NewNode("B")
    nc := NewNode("C")
    nd := NewNode("D")
    ne := NewNode("E")
    nf := NewNode("F")
    ng := NewNode("G")
    nh := NewNode("H")
    ni := NewNode("I")

    tree.Root = root
    root.left = nb
    root.right = nc
    root.left.left = nd
    root.left.right = ne
    root.right.left = nf
    root.right.right = ng
    root.left.left.left = nh
    root.left.left.right = ni
    tree.threaded()
    
    Convey("#PreOrder", t, func() {
        fmt.Print("\n")
        PreOrder(tree.Root, 0, 'M')
    })
    Convey("#printThreadedTree", t, func() {
        tree.printThreadedTree()
    })

}
