package engine

import "sync"

type ReadyQueue struct {
	mu    sync.Mutex
	queue []string
}

func NewReadyQueue() *ReadyQueue {
	return &ReadyQueue{}
}

func (q *ReadyQueue) Push(taskID string) {

	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, taskID)
}

func (q *ReadyQueue) Pop() (string, bool) {

	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		return "", false
	}

	task := q.queue[0]

	q.queue = q.queue[1:]

	return task, true
}