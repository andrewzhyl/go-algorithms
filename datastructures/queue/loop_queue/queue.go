package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func NewQueueWithCap(capacity int) *Queue {
	l := capacity + 1
	return &Queue{make([]interface{}, l), 0, 0, 0}
}

func NewQueue() *Queue {
	return &Queue{make([]interface{}, 10), 0, 0, 0}
}

func (q *Queue) GetCapaCity() int {
	return len(q.data) - 1
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.tail
}

func (q *Queue) GetSize() int {
	return q.size
}

func (q *Queue) Enqueue(x interface{}) {
	q.ResetSize(q.GetCapaCity() * 2)
	q.data[q.tail] = x
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *Queue) ResetSize(new_capacity int) {
	if q.front == (q.tail+1)%len(q.data) {
		new_data := make([]interface{}, new_capacity+1)
		for i := 0; i < q.GetSize(); i++ {
			new_data[i] = q.data[(i+q.front)%len(q.data)]
		}
		q.data = new_data
		q.front, q.tail = 0, q.size
		fmt.Println("queue resized to ", q.size)
	}
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Cannot dequeue from an empty queue.")
	}
	x := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.GetCapaCity()
	q.size--
	if q.size == q.GetCapaCity()/4 && q.GetCapaCity()/2 != 0 {
		q.ResetSize(q.GetCapaCity() / 2)
	}
	return x, nil
}

func (q *Queue) GetFront() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty.")
	}
	return q.data[q.front], nil
}

func (q *Queue) String() string {
	return fmt.Sprintf("cap: %d size: %d data: %v front: %d tail: %d", q.GetCapaCity(), q.GetSize(), q.data, q.front, q.tail)
}
