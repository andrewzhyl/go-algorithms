package sorts
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestShellSort(t *testing.T)  {
    items := []int{ 62,88,58,47,35,73,51,99,37,93}
    Convey("#ShellSort", t, func() {
        fmt.Printf("\n")
        ShellSort(items)
        fmt.Printf("希尔排序之后: ")
        fmt.Println(items)
    })

}