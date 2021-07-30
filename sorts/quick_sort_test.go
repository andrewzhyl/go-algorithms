package sorts
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestQuickSort(t *testing.T)  {
    items := []int{ 50,10,90,30,70,40,80,60, 88,20}
    Convey("#QuickSort", t, func() {
        QuickSort(items, 0, len(items) - 1)
        fmt.Printf("快速排序之后: ")
        fmt.Println(items)
    })
}