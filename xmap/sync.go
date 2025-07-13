package xmap

import "sync"

type SyncMap[Key any, Value any] struct {
	m sync.Map
}

func (m *SyncMap[Key, Value]) Get(key Key) (value Value, found bool) {
	res, found := m.m.Load(key)
	if !found {
		return
	}
	value, found = res.(Value)
	return
}

func (m *SyncMap[Key, Value]) Delete(key Key) {
	m.m.Delete(key)
}

func (m *SyncMap[Key, Value]) Add(key Key, value Value) {
	m.m.Store(key, value)
}

func (m *SyncMap[Key, Value]) ForEach(fn func(key Key, value Value) bool) {
	m.m.Range(func(key, value any) bool {
		return fn(key.(Key), value.(Value))
	})
}

func (m *SyncMap[Key, Value]) Has(key Key) bool {
	_, has := m.m.Load(key)
	return has
}
