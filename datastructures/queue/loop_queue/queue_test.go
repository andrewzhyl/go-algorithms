package queue

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	Convey("#GetCapaCity", t, func() {
		s := NewQueue()
		So(s.GetCapaCity(), ShouldEqual, 9)

		s = NewQueueWithCap(10)
		So(s.GetCapaCity(), ShouldEqual, 10)
	})

	Convey("#GetFront", t, func() {
		s := NewQueueWithCap(2)
		s.Enqueue(1)
		s.Enqueue(2)
		s.Enqueue(3)
		x, _ := s.GetFront()
		fmt.Println(x)
	})

	Convey("#Enqueue", t, func() {
		s := NewQueueWithCap(2)
		s.Enqueue(1)
		s.Enqueue(2)
		s.Enqueue(3)
		fmt.Println(s)
	})

	Convey("#Dequeue", t, func() {
		s := NewQueueWithCap(2)
		s.Enqueue(1)
		s.Enqueue(2)
		s.Enqueue(3)
		fmt.Println(s)
		x, _ := s.Dequeue()
		fmt.Println("Dequeue:", x)
		So(x, ShouldEqual, 1)
		fmt.Println(s)
	})

}
