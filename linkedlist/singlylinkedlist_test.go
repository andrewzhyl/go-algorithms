package linkedlist
    
import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestGetElem(t *testing.T)  {
    Convey("#GetElem", t, func() {
        list := InitList()
        list.ListInsert(1, "Andrew")
        _, elem := list.GetElem(1)
        if elem != "Andrew" {
            t.Log("TestGetElem test fail")
        }
    })
}

func TestListInsert(t *testing.T)  {
    list := InitList()

    Convey("参数 i=2 不在线性表范围内", t, func() {
        err := list.ListInsert(2, "Sunnie")
        So(err.Error(), ShouldEqual, "参数 i=2 不在线性表范围内")
    })

    Convey("参数 i=21 不在线性表范围内", t, func() {
        var i = 0
        for{
            list.ListInsert(i, "test")
            i++
            if(i <= 20){
                break
            }
        }
        err := list.ListInsert(21, "Sunnie")
        So(err.Error(), ShouldEqual,"参数 i=21 不在线性表范围内")
        list.ClearList()
    })
    
    Convey("#GetElem  success", t, func() {
        list.ListInsert(1, "Andrew")
        list.ListInsert(2, "Sunnie")
        _, elem := list.GetElem(2)
        So(elem, ShouldEqual, "Sunnie")
    })
}

func TestListDelete(t *testing.T)  {
    list := InitList()

    Convey("线性表为空", t, func() {
        err, _ := list.ListDelete(1)
        So(err.Error(), ShouldEqual, "线性表为空")
    })

    Convey("参数 i=2 不在线性表范围内", t, func() {
        list.ListInsert(1, "Andrew")
        err, _ := list.ListDelete(2)
        So(err.Error(), ShouldEqual,"参数 i=2 不在线性表范围内")
        list.ClearList()
    })
    
    Convey("#ListDelete  success", t, func() {
        list.ListInsert(1, "Andrew")
        _, elem := list.ListDelete(1)
        So(elem, ShouldEqual,"Andrew")
    })
}