package triemap

type Node struct {
	Value int
	nodes map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		Value: 0,
		nodes: make(map[rune]*Node),
	}
}

type TrieMap struct {
	Root *Node
	size int
}

func NewTrieMap() *TrieMap {
	return &TrieMap{
		Root: NewNode(),
		size: 0,
	}
}

// 获取大小
func (m *TrieMap) GetSize() int {
	return m.size
}

// 是否为空
func (m *TrieMap) IsEmpty() bool {
	return m.size == 0
}

// 插入操作
func (m *TrieMap) Add(key string, val int) *TrieMap {
	cur := m.Root
	for _, c := range []rune(key) {
		if _, ok := cur.nodes[c]; !ok {
			cur.nodes[c] = NewNode()
		}
		cur = cur.nodes[c]
	}
	cur.Value = val
	m.size++
	return m
}

func (m *TrieMap) GetNode(key string) *Node {
	cur := m.Root
	for _, c := range []rune(key) {
		if _, ok := cur.nodes[c]; !ok {
			return nil
		}
		cur = cur.nodes[c]
	}
	if cur.Value > 0 {
		return cur
	} else {
		return nil
	}
}

func (t *TrieMap) Contains(key string) bool {
	return t.GetNode(key) != nil
}

func (t *TrieMap) Get(key string) interface{} {
	node := t.GetNode(key)
	if node != nil {
		return node.Value
	}
	return nil
}

func (ll *TrieMap) Set(key string, val int) {
	node := ll.GetNode(key)
	if node == nil {
		panic(key + " doesn't exist!")
	}
	node.Value = val
}
