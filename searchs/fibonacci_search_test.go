package searches
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestFibonacci(t *testing.T)  {
    Convey("#Fibonacci", t, func() {
       items := Fibonacci(36)
       fmt.Printf("斐波那契数列: ")
       fmt.Println(items)
       So(len(items), ShouldEqual, 36)
    })
}

func TestFibonacciSearch(t *testing.T)  {
    Convey("#FibonacciSearch", t, func() {
        items := []int{0,16,24,35,47,59,62,73,88,99}
        index := FibonacciSearch(items, 59)
        fmt.Printf("斐波那契查找结果: ")
        fmt.Println(index)
        So(FibonacciSearch(items, 0), ShouldEqual, 0)
        So(FibonacciSearch(items, 59), ShouldEqual, 5)
        So(FibonacciSearch(items, 62), ShouldEqual, 6)
        So(FibonacciSearch(items, 73), ShouldEqual, 7)
        So(FibonacciSearch(items, 88), ShouldEqual, 8)
        So(FibonacciSearch(items, 99), ShouldEqual, 9)
    })
}