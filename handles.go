package generic

import (
	"log"
)

type Handle[T any] struct {
	slot int32
	gen  int32 // generation number!
}

type Manager[T any] struct {
	items     []T
	gens      []int32
	freeSlots []int32
	nextGen   int32
}

func MakeManager[T any]() *Manager[T] {
	return &Manager[T]{
		items:     make([]T, 0, 1024),
		gens:      make([]int32, 0, 1024),
		freeSlots: make([]int32, 0, 128),
	}
}

func (m *Manager[T]) Create() (itemPtr *T, handle Handle[T]) {
	m.nextGen++
	if len(m.freeSlots) > 0 {
		slot := Last(m.freeSlots)
		ShrinkTo(&m.freeSlots, len(m.freeSlots)-1)

		m.gens[slot] = m.nextGen
		handle.slot = slot
		handle.gen = m.nextGen
		itemPtr = &m.items[slot]
		return
	} else {
		Append(&m.gens, m.nextGen)
		itemPtr = AllocAppend(&m.items)
		handle.slot = int32(len(m.items) - 1)
		handle.gen = m.nextGen
		return
	}
}

func (m *Manager[T]) valid(handle Handle[T]) bool {
	return handle.slot >= 0 && handle.slot < int32(len(m.gens)) && m.gens[handle.slot] == handle.gen
}

func (m *Manager[T]) Delete(handle Handle[T]) {
	if !m.valid(handle) {
		log.Println("WARNING: handle already deleted!")
		return
	}
	Append(&m.freeSlots, handle.slot)
	Reset(&m.gens[handle.slot])
	Reset(&m.items[handle.slot])
}

func (m *Manager[T]) GetItem(handle Handle[T]) *T {
	if !m.valid(handle) {
		return nil
	} else {
		return &m.items[handle.slot]
	}
}

func (m *Manager[T]) Reset() {
	ResetSlice(&m.items)
	ResetSlice(&m.gens)
	ResetSlice(&m.freeSlots)
}
