package syncList

import (
	"container/list"
	"fmt"
	"sync"
)

type QueueData struct {
	Id string
	Value interface{}
}

type Queue struct {
	data *list.List
	lock sync.Mutex
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = list.New()
	return q
}

func (q *Queue) Push(v QueueData) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.data.PushFront(v)
}

func (q *Queue) Pop() QueueData {
	defer q.lock.Unlock()
	q.lock.Lock()
	iter := q.data.Back()
	v := iter.Value
	q.data.Remove(iter)
	return v.(QueueData)
}

func (q *Queue) Len() int {
	defer q.lock.Unlock()
	q.lock.Lock()
	return q.data.Len()
}

func (q *Queue) Delete(id string) {
	defer q.lock.Unlock()
	q.lock.Lock()
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		v := iter.Value
		if v.(QueueData).Id == id{
			q.data.Remove(iter)
		}
		fmt.Println("item:", iter.Value)
	}
}
