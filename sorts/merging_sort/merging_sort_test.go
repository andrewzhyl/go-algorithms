package sorts
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestMergingSort(t *testing.T)  {
    items := []int{ 50,10,90,30,70,40,80,60, 20}
    Convey("#MergingSort", t, func() {
        items = MergingSort(items)
        fmt.Printf("归并排序之后: ")
        fmt.Println(items)
    })
}