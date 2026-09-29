package reactive

import "testing"

func TestStoreSubscribeNotifyUnsubscribe(t *testing.T) {

	store := NewStore()

	received := make([]any, 0)

	unsubscribe := store.Subscribe("tasks", func(value any) {
		received = append(received, value)
	})

	store.Update("tasks", "one")
	store.Notify("tasks")

	if len(received) != 1 || received[0] != "one" {
		t.Fatalf("expected one notification with value one, got %#v", received)
	}

	store.Update("tasks", "two")
	store.Notify("tasks")

	if len(received) != 2 || received[1] != "two" {
		t.Fatalf("expected second notification with value two, got %#v", received)
	}

	unsubscribe()

	store.Update("tasks", "three")
	store.Notify("tasks")

	if len(received) != 2 {
		t.Fatalf("expected no notifications after unsubscribe, got %#v", received)
	}

	if store.Watchers("tasks") != 0 {
		t.Fatalf("expected no watchers after unsubscribe")
	}

}

func TestStoreRevisions(t *testing.T) {

	store := NewStore()

	if store.Revision("tasks") != 0 {
		t.Fatalf("expected initial revision to be 0")
	}

	store.Update("tasks", "a")

	if store.Revision("tasks") != 1 {
		t.Fatalf("expected revision to be 1")
	}

	store.Update("tasks", "b")

	if store.Revision("tasks") != 2 {
		t.Fatalf("expected revision to be 2")
	}

	if store.Revision("other") != 0 {
		t.Fatalf("expected unrelated revision to be 0")
	}

}

func TestStoreUnsubscribeIsIdempotent(t *testing.T) {

	store := NewStore()

	unsubscribe := store.Subscribe("tasks", func(value any) {})

	unsubscribe()
	unsubscribe()

	if store.Watchers("tasks") != 0 {
		t.Fatalf("expected idempotent unsubscribe")
	}

}

