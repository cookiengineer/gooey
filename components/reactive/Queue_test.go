package reactive

import "testing"

type fake_component struct {
	Name     string
	Rendered int
}

func TestQueueCoalesces(t *testing.T) {

	queue := NewQueue()

	a := &fake_component{Name: "a"}
	b := &fake_component{Name: "b"}

	queue.Push(a)
	queue.Push(a)
	queue.Push(b)
	queue.Push(a)

	if queue.Length() != 2 {
		t.Fatalf("expected 2 queued items, got %d", queue.Length())
	}

	drained := queue.Drain()

	if len(drained) != 2 || drained[0] != a || drained[1] != b {
		t.Fatalf("expected coalesced order [a b], got %#v", drained)
	}

	if queue.Length() != 0 {
		t.Fatalf("expected queue to be empty after drain")
	}

}

func TestQueueRemove(t *testing.T) {

	queue := NewQueue()

	a := &fake_component{Name: "a"}
	b := &fake_component{Name: "b"}

	queue.Push(a)
	queue.Push(b)
	queue.Remove(a)

	if queue.Length() != 1 {
		t.Fatalf("expected 1 queued item, got %d", queue.Length())
	}

	drained := queue.Drain()

	if len(drained) != 1 || drained[0] != b {
		t.Fatalf("expected only b, got %#v", drained)
	}

	queue.Push(a)
	queue.Remove(a)

	if queue.Length() != 0 {
		t.Fatalf("expected removed item to stay removed")
	}

}
