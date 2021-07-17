package linkedlist

import (
	"errors"
	"fmt"
)

const MAXSIZE = 20 //定义数组最大长度

type ElemType string

// 线性表: 顺序存储结构
type SqList struct {
	data   [MAXSIZE]ElemType // 数组存储数据元素，最大值为 MAXSIZE
	length int               // 线性表当前长度
}

//若线性表为空，返回 true, 否则返回 false
func InitList() *SqList {
	l := &SqList{}
	l.length = 0
	return l
}

//若线性表为空，返回 true, 否则返回 false
func (l *SqList) ListEmpty() bool {
	if l.length == 0 {
		return true
	}
	return false
}

//将线性表清空
func (l *SqList) ClearList() {
	l.data = [MAXSIZE]ElemType{}
	l.length = 0
}

// 将线性表 L 中的第 i 个未知元素值返回给 e
func (l *SqList) GetElem(i int) (error, ElemType) {
	if l.length == 0 || i < 1 || i > l.length {
		return errors.New("没有找到"), ""
	}
	e := l.data[i-1]
	return nil, e
}

// `LocateElem(L, e)`:  在线性表 L 中查找与给定值 e 相等的元素，如果查找成功，返回该元素在表中序号表示成功；否则，返回 0 表示失败
func (l *SqList) LocateElem(e ElemType) (error, int) {
	for i, a := range l.data {
		if a == e {
			return nil, i
		}
	}
	return errors.New("没有找到"), 0
}

// 在线性表 L 中的第 i 个位置插入新元素 e
func (l *SqList) ListInsert(i int, e ElemType) error {
	if l.length == MAXSIZE {
		return errors.New("顺序线性表已满")
	}

	if i < 1 || i > l.length+1 {
		return errors.New(fmt.Sprintf("参数 i=%v 不在线性表范围内", i))
	}

	if i <= l.length {
		for k := l.length - 1; k >= i-1; k-- {
			l.data[k+1] = l.data[k]
		}
	}

	l.data[i-1] = e // 将新元素插入
	l.length++
	return nil
}

// 删除线性表 L 中的第 i 个位置元素，并用 e 返回其值
func (l *SqList) ListDelete(i int) (err error, e ElemType) {
	if l.length == 0 { // 线性表为空
		return errors.New("线性表为空"), ""
	}

	if i < i || i > l.length {
		return errors.New(fmt.Sprintf("参数 i=%v 不在线性表范围内", i)), ""
	}

	e = l.data[i-1]
	if i < l.length {
		for k := i; k < l.length; k++ {
			l.data[k-1] = l.data[k]
		}
	}

	l.length--
	return nil, e
}

