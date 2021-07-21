package binarytree

import(
	"fmt"
	"sort"
)

type Node struct{
	val interface{}
	weight int
	parent *Node
	left *Node
	right *Node
}

// SortNodes 用于 Node 以 weight 排序 
type SortNodes []*Node

type Tree struct{
	Root *Node
}

func (sn SortNodes) Len() int {
	return len(sn)
}

func (sn SortNodes) Swap(i, j int) {
	sn[i], sn[j] = sn[j], sn[i]
}

func (sn SortNodes) Less(i, j int) bool {
	return sn[i].weight < sn[j].weight
}

// 创建哈夫曼树
func CreateHuffmanTree(leaves []*Node) *Node {
	if len(leaves) < 2 {
		return nil
	}
	sort.Stable(SortNodes(leaves))
	
	for len(leaves) > 1 {
		left, right := leaves[0], leaves[1]
		parentWeight := left.weight + right.weight
		parent := &Node{ val: ' ', left: left, right: right, weight: parentWeight }
		left.parent = parent
		right.parent = parent
		ls := leaves[2:]
		idx := sort.Search(len(ls), func(i int) bool { return ls[i].weight >= parentWeight })
		idx += 2
		
		// Insert
		copy(leaves[1:], leaves[2:idx])
		leaves[idx-1] = parent
		leaves = leaves[1:]
	}
	return leaves[0]
}

// Print the Huffman tree 
func PreOrderPrint(root *Node) {
	var traverse func(node *Node, ns int, ch rune)
	traverse = func(node *Node, ns int, ch rune) {
		if node == nil {
			return
		}
		for i := 0; i < ns; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%c:%c:%d\n", ch, node.val, node.weight)
		traverse(node.left, ns+2, 'L')
		traverse(node.right, ns+2, 'R')
	}
	traverse(root, 0, 'M')
}