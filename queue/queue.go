package queue

// An FIFO queue.
type Queue []interface{}

// Pushes the element into the queue
// 		e.g q.Push(123)
func (queue *Queue) Push(v int) {
	*queue = append(*queue, v)
}

// Pops element from head
func (queue *Queue) Pop() int {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head.(int)
}

// Returns if the queue is empty or not
func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}
