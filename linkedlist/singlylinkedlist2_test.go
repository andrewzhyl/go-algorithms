package linkedlist
    
import (
    // "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestAddFront(t *testing.T)  {
    Convey("#AddFront", t, func() {
        list := CreateList()
        list.AddFront(5)
        list.AddFront(6)
        list.AddFront(7)

        want := []interface{}{7, 6, 5}
        i := 0
        for cur := list.Head; cur.Next != nil; cur = cur.Next {
            So(cur.Val, ShouldEqual, want[i])
            i++
        }
    })
}

func TestAddBack(t *testing.T)  {
    Convey("#AddBack", t, func() {
        list := CreateList()
        list.AddBack(5)
        list.AddBack(6)
        list.AddBack(7)

        want := []interface{}{5, 6, 7}
        i := 0
        for cur := list.Head; cur.Next != nil; cur = cur.Next {
            So(cur.Val, ShouldEqual, want[i])
            i++
        }
    })
}

func TestDelFront(t *testing.T)  {
    Convey("#DelFront", t, func() {
        list := CreateList()
        _, err := list.DelFront()
        So(err.Error(), ShouldEqual, "链表为空，不能删除")
        list.AddBack(5)
        list.AddBack(6)
        list.AddBack(7)

        want := []interface{}{5, 6, 7}
        i := 0
        for (i < 3) {
            val, _ := list.DelFront()
            So(val, ShouldEqual, want[i])
            i++
        }
    })
}

func TestDelBack(t *testing.T)  {
    Convey("#DelBack", t, func() {
        list := CreateList()
        _, err := list.DelBack()
        So(err.Error(), ShouldEqual, "链表为空，不能删除")
        list.AddBack(15)
        list.AddBack(16)
        list.AddBack(17)
        So(list.Count(), ShouldEqual, 3)
        want := []interface{}{17, 16, 15}
        i := 0
        for (i < 3) {
            val, _ := list.DelBack()
            So(val, ShouldEqual, want[i])
            i++
        }
    })
}

func TestReverse(t *testing.T)  {
    Convey("#Reverse", t, func() {
        list := CreateList()
        list.AddBack(15)
        list.AddBack(16)
        list.AddBack(17)
        So(list.Count(), ShouldEqual, 3)
        list.Reverse()
        want := []interface{}{17, 16, 15}
        i := 0
        for (i < 3) {
            val, _ := list.DelFront()
            So(val, ShouldEqual, want[i])
            i++
        }
    })
}