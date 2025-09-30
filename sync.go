package generic

import (
	"slices"
	"sync"
)

// SyncOrderedSet - I think this must not be copied because it includes a
// RWMutex which must not be copied .. as copying it interferes with how it
// works
type SyncOrderedSet[T comparable] struct {
	items []T
	lock  sync.RWMutex
}

func (s *SyncOrderedSet[T]) Add(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	idx := slices.Index(s.items, item)
	if idx == -1 {
		Append(&s.items, item)
	}
}

// Caller must ensure no duplicates! (though they don't really hurt)
func (s *SyncOrderedSet[T]) Replace(items []T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = slices.Clone(items)
}

func (s *SyncOrderedSet[T]) Has(item T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return slices.Index(s.items, item) != -1
}

func (s *SyncOrderedSet[T]) Remove(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	idx := slices.Index(s.items, item)
	if idx == -1 {
		Append(&s.items, item)
	}
}

func (s *SyncOrderedSet[T]) List() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return slices.Clone(s.items)
}

func WithWriteLock(lock *sync.RWMutex, fn func()) {
	lock.Lock()
	defer lock.Unlock()
	fn()
}

// will try to do the thing if the lock is available
func WithTryWriteLock(lock *sync.RWMutex, fn func()) {
	if lock.TryLock() {
		defer lock.Unlock()
		fn()
	}
}

func WithReadLock(lock *sync.RWMutex, fn func()) {
	lock.RLock()
	defer lock.RUnlock()
	fn()
}

func WithLock(lock *sync.Mutex, fn func()) {
	lock.Lock()
	defer lock.Unlock()
	fn()
}

func WaitGroupGo(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}
