package segment_tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSegmentTree(t *testing.T) {
	Convey("#build", t, func() {
		nums := []Elem{-2, 0, 3, -5, 2, -1}
		tree := NewSegmentTree(nums, func(a, b Elem) Elem {
			return a.(int) + b.(int)
		})
		So("[-3,1,-4,-2,3,-3,-1,-2,0,nil,nil,-5,2,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil]", ShouldEqual, tree.String())
	})

	Convey("#query", t, func() {
		nums := []Elem{1, 2, 3, 4, 5}
		tree := NewSegmentTree(nums, func(a, b Elem) Elem {
			return a.(int) + b.(int)
		})
		So(6, ShouldEqual, tree.Query(0, 2))
		So(12, ShouldEqual, tree.Query(2, 4))
		So(15, ShouldEqual, tree.Query(0, 4))
	})
}
