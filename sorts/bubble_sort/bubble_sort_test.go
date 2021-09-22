package sorts
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestBubbleSort(t *testing.T)  {
    items := []int{ 62,88,58,47,35,73,51,99,37,93}
    // items := []int{ 35,37,47,51,62,58,73,88,93,99 }
    Convey("#BubbleSort", t, func() {
        BubbleSort(items)
        fmt.Printf("冒泡排序之后: ")
        fmt.Println(items)
    })

}