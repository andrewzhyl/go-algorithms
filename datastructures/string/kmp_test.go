package kmp
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestKmpTable(t *testing.T)  {
    Convey("#kmpTable", t, func() {
        next := kmpTable("ABCABD")
        // So(s.GetTop(), ShouldEqual, 6)
        fmt.Println(next)
    })
}