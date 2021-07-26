package sorts
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestInsertionSort(t *testing.T)  {
    items := []int{ 62,88,58,47,35,73,51,99,37,93}
    Convey("#InsertionSort", t, func() {
        InsertionSort(items)
        fmt.Printf("直接插入排序之后: ")
        fmt.Println(items)
    })

}