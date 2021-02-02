package main

import (
	"fmt"
	"container/list"
)

type intQueue struct {
	queue *list.List
}

func (iq *intQueue) enqueue(val int) {
	fmt.Println("Adding element", val, "in queue")
	iq.queue.PushBack(val)
}

func (iq *intQueue) dequeue() error {
	if iq.queue.Len() > 0 {
		ele := iq.queue.Front()
		iq.queue.Remove(ele)
	}
	return fmt.Errorf("Queue is empty")
}

func (iq *intQueue) peep() (int, error) {
	if iq.queue.Len() > 0 {
		if val, ok := iq.queue.Front().Value.(int); ok {
			return val, nil
		}
		return -1, fmt.Errorf("Peep Error: data type is incorrect") 
	}
	return -1, fmt.Errorf("Queue is empty")
}

func (iq *intQueue) size() int {
	return iq.queue.Len()
}

func (iq *intQueue) empty() bool {
	return iq.queue.Len() == 0
}

func main() {
	customQueue := &intQueue{
		queue: list.New(),
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
