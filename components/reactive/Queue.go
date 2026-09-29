package reactive

import "sync"

// Queue is a coalescing FIFO queue. Pushing an item that is already queued is a
// no-op, so multiple invalidations in one tick result in a single drain entry.
type Queue struct {
	mutex  sync.Mutex
	order  []any
	queued map[any]bool
}

// NewQueue creates an empty Queue.
func NewQueue() *Queue {

	queue := &Queue{
		order:  make([]any, 0),
		queued: make(map[any]bool),
	}

	return queue

}

// Push adds an item unless it is already queued.
func (queue *Queue) Push(item any) {

	if item == nil {
		return
	}

	queue.mutex.Lock()

	if is_comparable(item) == true && queue.queued[item] == true {
		queue.mutex.Unlock()
		return
	}

	queue.order = append(queue.order, item)

	if is_comparable(item) == true {
		queue.queued[item] = true
	}

	queue.mutex.Unlock()

}

// Remove drops an item from the queue, if present.
func (queue *Queue) Remove(item any) {

	if item == nil || is_comparable(item) == false {
		return
	}

	queue.mutex.Lock()

	if queue.queued[item] == true {

		delete(queue.queued, item)

		for i, queued := range queue.order {
			if queued == item {
				queue.order = append(queue.order[:i], queue.order[i+1:]...)
				break
			}
		}

	}

	queue.mutex.Unlock()

}

// Drain returns all queued items in order and empties the queue.
func (queue *Queue) Drain() []any {

	queue.mutex.Lock()

	result := queue.order
	queue.order = make([]any, 0)
	queue.queued = make(map[any]bool)

	queue.mutex.Unlock()

	return result

}

// Len returns the number of queued items.
func (queue *Queue) Length() int {

	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	return len(queue.order)

}

