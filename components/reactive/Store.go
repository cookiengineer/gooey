package reactive

import "sync"

// Store is a namespaced, observable key/value store.
//
// Watchers receive the original in-memory value, not a serialized copy, so
// callers may store typed pointers and type assert inside their callback.
type Store struct {
	mutex     sync.RWMutex
	values    map[string]any
	revisions map[string]uint64
	watchers  map[string]map[uint]func(value any)
	next_id   uint
}

// NewStore creates an empty Store.
func NewStore() *Store {

	store := &Store{
		values:    make(map[string]any),
		revisions: make(map[string]uint64),
		watchers:  make(map[string]map[uint]func(value any)),
	}

	return store

}

// Get returns the current value and whether it has been set.
func (store *Store) Get(name string) (any, bool) {

	store.mutex.RLock()
	defer store.mutex.RUnlock()

	value, ok := store.values[name]

	return value, ok

}

// Revision returns the current revision for a name. Unknown names are 0.
func (store *Store) Revision(name string) uint64 {

	store.mutex.RLock()
	defer store.mutex.RUnlock()

	return store.revisions[name]

}

// Set updates the in-memory value and bumps the revision. It does not notify
// watchers; call Notify afterwards once persistence (if any) has completed.
func (store *Store) Update(name string, value any) uint64 {

	store.mutex.Lock()

	store.values[name] = value
	store.revisions[name] = store.revisions[name] + 1
	revision := store.revisions[name]

	store.mutex.Unlock()

	return revision

}

// Notify invokes all watchers registered for name with the current value.
// Watchers are snapshotted before invocation so a callback may safely call
// Set, Notify or Subscribe without deadlocking.
func (store *Store) Notify(name string) {

	store.mutex.RLock()

	value, _ := store.values[name]
	callbacks := make([]func(value any), 0)

	if watchers, ok := store.watchers[name]; ok == true {
		for _, callback := range watchers {
			callbacks = append(callbacks, callback)
		}
	}

	store.mutex.RUnlock()

	for _, callback := range callbacks {
		callback(value)
	}

}

// Subscribe registers callback for name and returns an unsubscribe function.
func (store *Store) Subscribe(name string, callback func(value any)) func() {

	store.mutex.Lock()

	store.next_id++
	id := store.next_id

	if _, ok := store.watchers[name]; ok == false {
		store.watchers[name] = make(map[uint]func(value any))
	}

	store.watchers[name][id] = callback

	store.mutex.Unlock()

	var once sync.Once

	return func() {

		once.Do(func() {

			store.mutex.Lock()

			if watchers, ok := store.watchers[name]; ok == true {
				delete(watchers, id)
			}

			store.mutex.Unlock()

		})

	}

}

// Watchers returns the number of active watchers for a name.
func (store *Store) Watchers(name string) int {

	store.mutex.RLock()
	defer store.mutex.RUnlock()

	return len(store.watchers[name])

}
