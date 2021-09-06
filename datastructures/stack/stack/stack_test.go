package stack
    
import (
    // "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestPush(t *testing.T)  {
    Convey("#Push", t, func() {
        s := NewStack()
        s.Push(5)
        s.Push(6)
        So(s.GetTop(), ShouldEqual, 6)
    })
}

func TestPop(t *testing.T)  {
    Convey("#Pop", t, func() {
        s := NewStack()
        s.Push(5)
        s.Push(6)
        So(s.Pop(), ShouldEqual, 6)
        s.Push(7)
        So(s.Pop(), ShouldEqual, 7)
        s.Push(7)
        s.Push(7)
        s.Push(7)
        So(s.Len(), ShouldEqual, 4)
    })
}