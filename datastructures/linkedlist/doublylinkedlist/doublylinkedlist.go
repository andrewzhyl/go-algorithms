package linkedlist

import (
	"errors"
	"fmt"
)

const MAXSIZE = 20 //定义数组最大长度

type Node struct {
	val interface{}
	prev *Node
	next *Node
}

// DoubleLinkedList structure with length of the list and its head
type DoubleLinkedList struct {
	head   *Node
}

func CreateList() *DoubleLinkedList{
	return &DoubleLinkedList{}
}

func CeateNode(val interface{}) *Node {
	note := &Node{ }
	note.val = val
	note.prev = nil
	note.next = nil
	return note
}

func (ll *DoubleLinkedList) AddFront(val interface{}){
	elem := CeateNode(val)
	if ll.head != nil{
		ll.head.prev = elem
		elem.next = ll.head
	}
	ll.head = elem
}

func (ll *DoubleLinkedList) AddBack(val interface{}){
	elem := CeateNode(val)
	if ll.head == nil{
		ll.head = elem
	}else{
		current := ll.head
		for(current.next != nil){
			current = current.next
		}
		current.next = elem
		elem.prev = current
	}
}

func (ll *DoubleLinkedList) DelFront() (interface{}, error){
	if ll.head == nil{
		return nil, errors.New("链表为空，不能删除")
	}
	current := ll.head
	ll.head = current.next
	if ll.head != nil {
		ll.head.prev = nil
	}
	return  current.val, nil
}

func (ll *DoubleLinkedList) DelBack() (interface{}, error){
	if ll.head == nil{
		return nil, errors.New("链表为空，不能删除")
	}

	if ll.head.next == nil {
		return ll.DelFront()
	}

	current := ll.head
	for(current.next.next != nil){
		current = current.next
	}
	retval := current.next.val
	current.next = nil
	return  retval, nil
}

func (ll *DoubleLinkedList) Count() int {
	var ctr int = 0
	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}
	return ctr
}

func (ll *DoubleLinkedList) Reverse() error{
	if ll.head == nil{
		return errors.New("链表为空，不能反转")
	}
	var prev, next *Node
	cur := ll.head

	for cur != nil{
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}
	ll.head = prev
	return nil
}

func  (ll *DoubleLinkedList) Display() {
	for current := ll.head; current.next != nil; current = current.next {
		fmt.Print(current.val, " ")
	}
	fmt.Print("\n")
}