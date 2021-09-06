package linkedlist

import (
	"errors"
	"fmt"
)

const MAXSIZE = 20 //定义数组最大长度

type Node struct {
	Val interface{}
	Next *Node
}

// SingleLinkedList structure with length of the list and its head
type SingleLinkedList struct {
	length int
	Head   *Node
}

func CreateList() *SingleLinkedList{
	return &SingleLinkedList{}
}

func CeateNode(val interface{}) *Node {
	return &Node{ val, nil }
}

func (ll *SingleLinkedList) AddFront(val interface{}){
	elem := CeateNode(val)
	elem.Next = ll.Head
	ll.Head = elem
	ll.length++
}

func (ll *SingleLinkedList) AddBack(val interface{}){
	elem := CeateNode(val)
	if ll.Head == nil{
		ll.Head = elem
	}else{
		current := ll.Head
		for(current.Next != nil){
			current = current.Next
		}
		current.Next = elem
	}
	ll.length++
}

func (ll *SingleLinkedList) DelFront() (interface{}, error){
	if ll.Head == nil{
		return nil, errors.New("链表为空，不能删除")
	}
	current := ll.Head
	ll.Head = current.Next
	ll.length--
	return  current.Val, nil
}

func (ll *SingleLinkedList) DelBack() (interface{}, error){
	if ll.Head == nil{
		return nil, errors.New("链表为空，不能删除")
	}
	var prev *Node
	current := ll.Head
	for(current.Next != nil){
		prev = current
		current = current.Next
	}

	if prev != nil {
		prev.Next = nil
	}else{
		ll.Head = nil
	}
	ll.length--
	return  current.Val, nil
}

func (ll *SingleLinkedList) Count() int {
	return ll.length
}

func (ll *SingleLinkedList) Reverse() error{
	if ll.Head == nil{
		return errors.New("链表为空，不能反转")
	}
	var prev, next *Node
	cur := ll.Head

	for cur != nil{
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	ll.Head = prev
	return nil
}

func  (ll *SingleLinkedList) Display() {
	for current := ll.Head; current.Next != nil; current = current.Next {
		fmt.Print(current.Val, " ")
	}
	fmt.Print("\n")
}