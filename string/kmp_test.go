package kmp
    
import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestGetNext(t *testing.T)  {
    Convey("#GetNext", t, func() {
        next := GetNext("ababaaaba")
        // So(s.GetTop(), ShouldEqual, 6)
        fmt.Println(next)
    })
}