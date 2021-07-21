package binarytree
    
import (
    "fmt"
    // "sort"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestCreateHuffmanTree(t *testing.T)  {
    Convey("#CreateHuffmanTree", t, func() {
       leaves := []*Node{ // From "this is an example of a huffman tree"
        {val: 'A', weight: 27},
        {val: 'B', weight: 8},
        {val: 'C', weight: 15},
        {val: 'D', weight: 15},
        {val: 'E', weight: 30},
        {val: 'F', weight: 5 },
        }
        fmt.Printf("\n")
        root := CreateHuffmanTree(leaves)
        PreOrderPrint(root)
    })
}