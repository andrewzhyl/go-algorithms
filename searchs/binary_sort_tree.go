package searches

import(
	"fmt"
)

type Node struct {
    val int
    left  *Node 
    right *Node 
}

// 二叉查找树
type BinarySortTree struct {
	root *Node
}

func NewBinarySortTree(val int) *BinarySortTree {
	node := &Node{ val, nil, nil }
	return &BinarySortTree{ node }
}

func NewNode(val int) *Node {
	return &Node{ val, nil, nil }
}

// 插入操作
func (bst *BinarySortTree) Insert(val int) *BinarySortTree{
	if bst.root == nil {
		bst.root = NewNode(val)
	} else{
		bst.root.Insert(val)
	}
	return bst
}

// 递归插入
func (n *Node) Insert(val int) *Node{
	if( n == nil){
		return NewNode(val)
	} 
	if val < n.val {
		n.left = n.left.Insert(val)
	} else if val > n.val {
		n.right = n.right.Insert(val)
	}
	return n
}

// 非递归插入
func (n *Node) Insert2(val int){
	if( n == nil){
		return 
	} 
	var terminalNode *Node
	for n != nil {
		terminalNode = n
		if val < n.val {
			n = n.left
		} else {
			n = n.right
		}
	}

	if val < terminalNode.val {
		terminalNode.left = NewNode(val)
	} else {
		terminalNode.right = NewNode(val)
	}
	return 
}

// 二叉树查找
func (t *BinarySortTree) Search(val int) *Node  {
	if t == nil || t.root == nil {
		return nil
	}
	return t.root.Search(val)
}

// 查找
func (n *Node) Search(key int) *Node {
	if n == nil {
		return nil
	}

	if key == n.val{
		return n
	}
	if key < n.val {
		return n.left.Search(key)
	} 
	return  n.right.Search(key)
}

// 二叉树删除
func (t *BinarySortTree) Delete(key int) {
	// 如果为空，直接返回
	if t == nil || t.root == nil {
		return 
	}
	t.root.Delete(key)
}

// 二叉树删除
func (node *Node) Delete(key int) *Node {
	if node == nil {
		return node
	}	

	// 如果删除的值小于 node 的值，则递归的遍历 node 的左子树
	// 删除的结点位于左子树
	if key < node.val {
		node.left = node.left.Delete(key)

	// 如果删除的值大于 node 的值，则递归的遍历 node 的右子树
	// 删除的结点位于右子树
	} else if key > node.val {
		node.right = node.right.Delete(key)

	// 如果值相等，则删除结点
	}else {
		// 一、被删除的结点是叶子结点：直接从二叉排序树当中移除即可
		// 二、被删除的结点有 1 个子结点，子结点需要连接到父亲结点
		if node.left == nil {
			node = node.right
			return node
		} else if node.right == nil {
			node = node.left
			return node
		}

		// 三、被删除的结点有左右子结点都不为空
		
		// 获取删除结点中序遍历的直接前驱结点
		temp := node.left.GetMaxNode()

		// 将删除结点的值替换为直接后继结点的值
		node.val = temp.val

		// 删除直接后继结点
		node.left = node.left.Delete(temp.val)

		// 或者，获取删除结点中序遍历的直接后区结点
		// temp := node.right.GetMinNode()
		// node.val = temp.val
		// node.right = node.right.Delete(temp.val)
		
	}
	return node
}

// 获取最大结点
func (node *Node) GetMaxNode() *Node {
	if node == nil {
		return nil
	}
	for {
		if node.right == nil{
			return node
		}
		node = node.right
	}
}

// 获取最小结点
func (node *Node) GetMinNode() *Node {
	if node == nil {
		return nil
	}
	for {
		if node.left == nil{
			return node
		}
		node = node.left
	}
}

// 前序遍历打印
func (n *Node) PreOrderTraverse(){
	if n == nil {
		return 
	}
	fmt.Print(" -> ", n.val)
	if n.left != nil{
		n.left.PreOrderTraverse()
	}
	if n.right != nil{
		n.right.PreOrderTraverse()
	}
}

// 中序遍历打印
func (n *Node) InOrderTraverse(){
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

// 后序遍历打印
func (n *Node) PostOrderTraverse(){
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

// Print the Huffman tree 
func PreOrderPrint(root *Node) {
	fmt.Printf("\n")
	var traverse func(node *Node, ns int, ch rune)
	traverse = func(node *Node, ns int, ch rune) {
		if node == nil {
			return
		}
		for i := 0; i < ns; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%c:%d\n", ch, node.val)
		traverse(node.left, ns+2, 'L')
		traverse(node.right, ns+2, 'R')
	}
	traverse(root, 0, 'M')
}