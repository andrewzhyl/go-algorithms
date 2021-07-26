package sorts
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestSelectionSort(t *testing.T)  {
    items := []int{ 62,88,58,47,35,73,51,99,37,93}
    Convey("#SelectionSort", t, func() {
        SelectionSort(items)
        fmt.Printf("简单选择排序之后: ")
        fmt.Println(items)
    })

}