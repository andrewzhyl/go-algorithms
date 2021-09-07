package trie

type Node struct {
	isWordEnd bool
	nodes     map[rune]*Node
}

func NewNode(isWordEnd bool) *Node {
	return &Node{isWordEnd, make(map[rune]*Node)}
}

type Trie struct {
	root *Node
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(false),
		size: 0,
	}
}

func (t *Trie) GetSize() int {
	return t.size
}

func (s *Trie) IsEmpty() bool {
	return s.size == 0
}

// 向 trie 中添加一个单词
func (t *Trie) Add(word string) {
	if t.Contains(word) {
		return
	}
	cur := t.root
	for _, c := range []rune(word) {
		if _, ok := cur.nodes[c]; !ok {
			cur.nodes[c] = NewNode(false)
		}
		cur = cur.nodes[c]
	}
	if !cur.isWordEnd {
		cur.isWordEnd = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, c := range []rune(word) {
		if _, ok := cur.nodes[c]; !ok {
			return false
		}
		cur = cur.nodes[c]
	}
	return cur.isWordEnd
}

func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for _, c := range []rune(prefix) {
		if _, ok := cur.nodes[c]; !ok {
			return false
		}
		cur = cur.nodes[c]
	}
	return true
}
