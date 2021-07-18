package queue

import(
	"fmt"
)

type Queue struct {
	data []int
	length int
}

func NewQueue(cap int) *Queue {
	return &Queue{ data: make([]int, 0, cap), length: 0 }
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	if(q.length == 0){
		return nil
	}
	n := q.data[0]
	q.length--
	q.data = q.data[1:]
	return n
}

func (q *Queue) Len() interface{} {
	return q.length
}

func (q *Queue) GetTop() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.data[0]
}

func (q *Queue) String() string {
    return fmt.Sprint(q.data)
}