//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/animations"
import "github.com/cookiengineer/gooey/components/interfaces"
import "github.com/cookiengineer/gooey/components/reactive"
import "sync"

// Scheduler coalesces invalidated components and renders them once per
// animation frame. It is safe to schedule from goroutines.
type Scheduler struct {
	mutex   sync.Mutex
	queue   *reactive.Queue
	running bool
	started bool
	pending bool
	frame   uint
}

// NewScheduler creates a stopped scheduler.
func NewScheduler() *Scheduler {

	return &Scheduler{
		queue: reactive.NewQueue(),
	}

}

// Schedule queues a component for the next flush.
func (scheduler *Scheduler) Schedule(component interfaces.Component) {

	scheduler.queue.Push(component)

	scheduler.mutex.Lock()
	started := scheduler.started
	scheduler.mutex.Unlock()

	if started == true {
		scheduler.RequestFrame()
	}

}

// Unschedule removes a component that is no longer mounted.
func (scheduler *Scheduler) Unschedule(component interfaces.Component) {
	scheduler.queue.Remove(component)
}

// Flush renders every queued component. It is safe to call synchronously, for
// example from tests, and is a no-op while a flush is already running.
func (scheduler *Scheduler) Flush() {

	scheduler.mutex.Lock()

	if scheduler.running == true {
		scheduler.mutex.Unlock()
		return
	}

	scheduler.running = true
	scheduler.mutex.Unlock()

	items := scheduler.queue.Drain()

	for _, item := range items {

		component, ok := item.(interfaces.Component)

		if ok == true && component != nil {
			component.Render()
		}

	}

	scheduler.mutex.Lock()
	scheduler.running = false
	scheduler.mutex.Unlock()

}

// Start enables automatic frame scheduling.
func (scheduler *Scheduler) Start() {

	scheduler.mutex.Lock()
	scheduler.started = true
	scheduler.mutex.Unlock()

}

// Stop disables automatic frame scheduling and cancels a pending frame.
func (scheduler *Scheduler) Stop() {

	scheduler.mutex.Lock()

	scheduler.started = false
	frame := scheduler.frame
	pending := scheduler.pending

	scheduler.pending = false
	scheduler.frame = 0

	scheduler.mutex.Unlock()

	if pending == true {
		animations.CancelAnimationFrame(frame)
	}

}

// Pending returns the number of queued components.
func (scheduler *Scheduler) Pending() int {
	return scheduler.queue.Length()
}

func (scheduler *Scheduler) RequestFrame() {

	scheduler.mutex.Lock()
	defer scheduler.mutex.Unlock()

	if scheduler.pending == true {
		return
	}

	scheduler.pending = true

	scheduler.frame = animations.RequestAnimationFrame(func(timestamp float64) {

		scheduler.mutex.Lock()
		scheduler.pending = false
		scheduler.frame = 0
		scheduler.mutex.Unlock()

		scheduler.Flush()

	})

}
