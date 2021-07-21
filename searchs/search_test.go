package searches
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearch(t *testing.T)  {
    Convey("#BinarySearch", t, func() {
       items := []int{0,1,16,24,35,47,59,62,73,88,99}
       target := BinarySearch(items, 62)
       fmt.Printf("二分查找索引是: ")
       fmt.Print(target)
       So(target, ShouldEqual, 7)
    })
}