package queue

import(
	"fmt"
)

type Node struct {
	val interface{}
	next *Node
}

type Queue struct {
	start, end *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{ nil, nil, 0 }
}

func (q *Queue) Enqueue(val interface{}) bool {
	elem := &Node{ val, nil }
	if(q.length == 0){
		q.start = elem
		q.end = elem
	} else{
		q.end.next = elem
		q.end = elem
	}
	q.length++
	return true
}

func (q *Queue) Dequeue() interface{} {
	if(q.length == 0){
		return nil
	}
	elem := q.start
	if(q.length == 1){
		q.start = nil
		q.end = nil
	}else{
		q.start = q.start.next
	}
	q.length--
	return elem.val
}

func (q *Queue) Len() interface{} {
	return q.length
}

func (q *Queue) GetTop() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.start.val
}

func (q *Queue) String() string {
	if(q.length == 0){
		return ""
	}
	cur := q.start
	s :=  fmt.Sprintf("%d", cur.val)
	for cur.next != nil {
		 cur = cur.next
		 s = fmt.Sprintf("%s %d", s, cur.val)
	}
    return fmt.Sprintf("[%s]", s)
}