package main

import (
	"fmt"
	"sync"
)

type intQueue struct {
	queue []int
	lock  sync.RWMutex
}

func (iq *intQueue) enqueue(val int) {
	fmt.Println("Adding element", val, "in queue")
	iq.lock.Lock()
	defer iq.lock.Unlock()
	iq.queue = append(iq.queue, val)
}

func (iq *intQueue) dequeue() error {
	if len(iq.queue) > 0 {
		iq.lock.Lock()
		defer iq.lock.Unlock()
		iq.queue = iq.queue[1:]
		return nil
	}
	return fmt.Errorf("Queue is empty")
}

func (iq *intQueue) peep() (int, error) {
	if len(iq.queue) > 0 {
		iq.lock.Lock()
		defer iq.lock.Unlock()
		return iq.queue[0], nil
	}
	return -1, fmt.Errorf("Queue is empty")
}

func (iq *intQueue) size() int {
	return len(iq.queue)
}

func (iq *intQueue) empty() bool {
	return len(iq.queue) == 0
}

func main() {
	customQueue := &intQueue{
		queue: make([]int, 0),
	}
	err := customQueue.dequeue()
	if err != nil {
		fmt.Println(err)
	}
	customQueue.enqueue(2)
	customQueue.enqueue(3)
	customQueue.enqueue(1)
	customQueue.enqueue(3)
	fmt.Println("Size of queue is:", customQueue.size())
	fmt.Println("Removing one element")
	customQueue.dequeue()
	val, _ := customQueue.peep()
	fmt.Println("front element:", val)
}
