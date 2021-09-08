package unionfind

type UnionFind struct {
	parent []int
	rank   []int // rank[i] 表示以 i 为根的树的层数
}

func NewUnionFind(N int) *UnionFind {
	uf := new(UnionFind)
	uf.parent = make([]int, N)
	uf.rank = make([]int, N)
	for i := 0; i < N; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (uf *UnionFind) GetSize() int {
	return len(uf.parent)
}

// 查找过程，查找元素 p 所对应的集合编号，所谓的根节点
// O(h) 复杂度，h 为树的高度
func (uf *UnionFind) Find(p int) int {
	if p < 0 || p >= len(uf.parent) {
		panic("p is out of bound.")
	}
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}

func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)

	if proot == qroot {
		return
	}

	// 根据两个元素所在树的 rank 层数不同判断合并方向
	// 将 rank 低的集合合并到 rank 高的集合上
	if uf.rank[proot] < uf.rank[qroot] {
		uf.parent[proot] = qroot
	} else if uf.rank[qroot] < uf.rank[proot] {
		uf.parent[qroot] = proot
	} else { //  uf.rank[qroot] <= uf.rank[proot] {
		uf.parent[qroot] = proot
		uf.rank[proot]++
	}
}
