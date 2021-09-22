package avltree

import (
	"fmt"
	"math"
	"strings"
)

type K interface{}
type V interface{}

type Node struct {
	key    K
	Value  V
	left   *Node
	right  *Node
	height int
}

func NewNode(key K, value V) *Node {
	return &Node{key, value, nil, nil, 1}
}

// 二叉查找树
type Tree struct {
	Root *Node
	size int
}

func NewTree() *Tree {
	return &Tree{nil, 0}
}

// 获取大小
func (t *Tree) GetSize() int {
	return t.size
}

// 是否为空
func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

// 获得节点的平衡因子
func getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return getHeight(node.left) - getHeight(node.right)
}

// 对节点 进行向右旋转操作，返回旋转后新的根节点 x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2

func rightRotate(y *Node) *Node {
	x := y.left
	t3 := x.right

	// 向右旋转
	x.right = y
	y.left = t3

	// 更新节点 height 值
	updateHeight(y)
	updateHeight(x)
	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点 x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func leftRotate(y *Node) *Node {
	x := y.right
	t2 := x.left

	// 向左旋转
	x.left = y
	y.right = t2

	// 更新节点 height 值
	updateHeight(y)
	updateHeight(x)
	return x
}

// 更新节点高度
func updateHeight(node *Node) {
	if node == nil {
		return
	}
	node.height = maxHeight(node.left, node.right) + 1
}

// 获取两个节点最大的高度
func maxHeight(a, b *Node) int {
	return int(math.Max(float64(getHeight(a)), float64(getHeight(b))))
}

// 插入操作
func (t *Tree) Add(key, val interface{}) *Tree {
	t.Root = add(t, t.Root, key, val)
	return t
}

// 递归插入
func add(t *Tree, node *Node, key, val interface{}) *Node {
	if node == nil {
		t.size++
		return NewNode(key, val)
	}
	if compare(key, node.key) < 0 {
		node.left = add(t, node.left, key, val)
	} else if compare(key, node.key) > 0 {
		node.right = add(t, node.right, key, val)
	} else {
		node.Value = val
	}
	// 更新高度
	node.height = 1 + maxHeight(node.left, node.right)

	// 计算平衡因子
	balanceFactor := getBalanceFactor(node)
	if balanceFactor > 1 {
		// fmt.Println("unbalanced:", balanceFactor)
	}

	// 平衡维护

	//LL
	if balanceFactor > 1 && getBalanceFactor(node.left) >= 0 {
		return rightRotate(node)
	}
	//RR
	if balanceFactor < -1 && getBalanceFactor(node.right) <= 0 {
		return leftRotate(node)
	}

	//LR
	if balanceFactor > 1 && getBalanceFactor(node.left) < 0 {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	//RL
	if balanceFactor < -1 && getBalanceFactor(node.right) > 0 {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

// 是否是二分搜索树，根据中序便利判断顺序是否从小到大排列
func (t *Tree) isBST() bool {
	keys := make([]K, 0)
	t.Root.InOrder(&keys)
	for i := 1; i < len(keys); i++ {
		if compare(keys[i-1], keys[i]) > 0 {
			return false
		}
	}
	return true
}

// 中序遍历，将 key 放入 keys
func (n *Node) InOrder(keys *[]K) {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.InOrder(keys)
	}
	*keys = append(*keys, n.key)
	if n.right != nil {
		n.right.InOrder(keys)
	}
}

// 是否是平衡二叉树
func (t *Tree) isBalanced() bool {
	return isBalanced(t.Root)
}

func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	balancefactor := getBalanceFactor(node)
	if balancefactor > 1 {
		return false
	}
	return isBalanced(node.left) && isBalanced(node.right)
}

// 二叉树查找
func (t *Tree) GetNode(key interface{}) *Node {
	if t == nil || t.Root == nil {
		return nil
	}
	return t.Root.getNode(key)
}

// 查找
func (n *Node) getNode(key interface{}) *Node {
	if n == nil {
		return nil
	}

	if compare(key, n.key) == 0 {
		return n
	}
	if compare(key, n.key) < 0 {
		return n.left.getNode(key)
	}
	return n.right.getNode(key)
}

func (t *Tree) Contains(key interface{}) bool {
	return t.GetNode(key) != nil
}

func (ll *Tree) Get(key interface{}) interface{} {
	node := ll.GetNode(key)
	if node != nil {
		return node.Value
	}
	return nil
}

func (ll *Tree) Set(key, val interface{}) {
	node := ll.GetNode(key)
	if node == nil {
		panic(key.(string) + " doesn't exist!")
	}
	node.Value = val
}

// 二叉树删除
func (t *Tree) Delete(key interface{}) interface{} {
	node := t.GetNode(key)
	if node != nil {
		t.Root = delete(t, t.Root, key)
		return node.Value
	}
	return nil
}

// 二叉树删除
func delete(t *Tree, node *Node, key interface{}) *Node {
	if node == nil {
		// 如果是空树，直接返回
		return nil
	}

	var retNode *Node

	// 如果删除的值小于 node 的值，则递归的遍历 node 的左子树
	// 删除的结点位于左子树
	if compare(key, node.key) < 0 {
		// 从左子树开始删除
		node.left = delete(t, node.left, key)
		// return node
		retNode = node

		// 如果删除的值大于 node 的值，则递归的遍历 node 的右子树
		// 删除的结点位于右子树
	} else if compare(key, node.key) > 0 {

		// 从右子树开始删除
		node.right = delete(t, node.right, key)
		// return node
		retNode = node

		// 如果值相等，则删除结点
	} else {

		// 一、被删除的结点是叶子结点：直接从二叉排序树当中移除即可
		// 二、被删除的结点有 1 个子结点，子结点需要连接到父亲结点

		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			// return rightNode
			t.size--
			retNode = rightNode

			// 待删除节点右子树为空的情况
		} else if node.right == nil {
			leftNode := node.left
			node.left = nil
			t.size--
			// return leftNode
			retNode = leftNode

			// 三、被删除的结点有左右子结点都不为空
		} else {

			// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
			// 用这个节点顶替待删除节点的位置
			temp := node.right.GetMinNode()
			temp.right = delete(t, node.right, temp.key)
			temp.left = node.left

			node.right = nil
			node.left = nil

			retNode = temp
		}
	}
	if retNode == nil {
		return nil
	}

	updateHeight(retNode)

	// 计算平衡因子
	balanceFactor := getBalanceFactor(retNode)

	// 平衡维护
	// LL
	if balanceFactor > 1 && getBalanceFactor(retNode.left) >= 0 {
		return rightRotate(retNode)
	}

	// RR
	if balanceFactor < -1 && getBalanceFactor(retNode.right) <= 0 {
		return leftRotate(retNode)
	}

	// LR
	if balanceFactor > 1 && getBalanceFactor(retNode.left) < 0 {
		retNode.left = leftRotate(retNode.left)
		return rightRotate(retNode)
	}

	// RL
	if balanceFactor < -1 && getBalanceFactor(retNode.right) > 0 {
		retNode.right = rightRotate(retNode.right)
		return leftRotate(retNode)
	}

	return retNode
}

// 获取最大结点
func (node *Node) GetMaxNode() *Node {
	if node == nil {
		return nil
	}
	for {
		if node.right == nil {
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
		if node.left == nil {
			return node
		}
		node = node.left
	}
}

// 前序遍历打印
func (n *Node) PreOrderTraverse() {
	if n == nil {
		return
	}
	fmt.Printf(" -> %v:%d \n", n.key, n.Value)
	if n.left != nil {
		n.left.PreOrderTraverse()
	}
	if n.right != nil {
		n.right.PreOrderTraverse()
	}
}

// 中序遍历打印
func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.InOrderTraverse()
	}
	fmt.Printf(" -> %v:%d \n", n.key, n.Value)
	if n.right != nil {
		n.right.InOrderTraverse()
	}
}

func compare(a, b interface{}) int {
	switch a.(type) {
	case string:
		return strings.Compare(a.(string), b.(string))
	case int32:
		return func(a, b int32) int {
			if a == b {
				return 0
			}
			if a < b {
				return -1
			}
			return +1
		}(a.(int32), b.(int32))
	default:
		return func(a, b int) int {
			if a == b {
				return 0
			}
			if a < b {
				return -1
			}
			return +1
		}(a.(int), b.(int))
	}
}
