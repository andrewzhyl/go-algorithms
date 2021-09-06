package maxheap

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHeap(t *testing.T) {
	Convey("#Add", t, func() {
		n := 1000000
		heap := NewMaxHeap()
		for i := 0; i < n; i++ {
			heap.add(Elem(rand.Int()))
		}

		arr := make([]Elem, n)
		for i := 0; i < n; i++ {
			arr[i] = heap.ExceptMax()
		}

		for i := 1; i < n; i++ {
			So(arr[i-1].(int), ShouldBeGreaterThan, arr[i].(int))
		}
	})
}
