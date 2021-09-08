package unionfind

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func uftest(uf *UnionFind, m int) time.Duration {

	size := uf.GetSize()
	now := time.Now()

	for i := 0; i < m; i++ {
		uf.Union(rand.Intn(size), rand.Intn(size))
	}

	for i := 0; i < m; i++ {
		uf.IsConnected(rand.Intn(size), rand.Intn(size))
	}
	return time.Since(now)
}

func TestUnionfind(t *testing.T) {

	Convey("#TrieMap", t, func() {
		size := 10000000
		m := 10000000

		uf := NewUnionFind(size)
		fmt.Println("UnionFind : ", uftest(uf, m))

	})
}
