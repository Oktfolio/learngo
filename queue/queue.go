package main

import "fmt"

func main() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.isEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.isEmpty())
}

type Queue []int

func (queue *Queue) Push(v int) {
	*queue = append(*queue, v)
}

func (queue *Queue) Pop() int {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head
}

func (queue *Queue) isEmpty() bool {
	return len(*queue) == 0
}
