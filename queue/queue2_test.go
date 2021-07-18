package queue
    
import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestEnqueue(t *testing.T)  {
    Convey("#Enqueue", t, func() {
        s := NewQueue(5)
        s.Enqueue(5)
        s.Enqueue(6)
        s.Enqueue(7)
        So(s.GetTop(), ShouldEqual, 5)
        So(s.Len(), ShouldEqual, 3)
        n := s.Dequeue()
        So(n, ShouldEqual, 5)
        So(s.GetTop(), ShouldEqual, 6)
        So(s.Len(), ShouldEqual, 2)
        So("[6 7]", ShouldEqual, s.String())
    })
}

func TestDequeue(t *testing.T)  {
    Convey("#Dequeue", t, func() {
        s := NewQueue(2)
        s.Enqueue(21)
        s.Enqueue(22)
        s.Enqueue(23)
        n := s.Dequeue()
        So(n, ShouldEqual, 21)
        So(s.GetTop(), ShouldEqual, 22)
        So(s.Len(), ShouldEqual, 2)
        So("[22 23]", ShouldEqual, s.String())
    })
}