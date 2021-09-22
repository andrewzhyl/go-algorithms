package sorts
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestHeapSort(t *testing.T)  {
    items := []int{ 50,10,90,30,70,40,80,60,20, 88}
    Convey("#HeapSort", t, func() {
        HeapSort(items)
        fmt.Printf("堆排序之后: ")
        fmt.Println(items)
    })
}