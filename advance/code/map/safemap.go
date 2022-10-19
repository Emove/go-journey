package _map

import (
	"sync"
)

type rwMap struct {
	rw sync.RWMutex
	m  map[interface{}]int
}

type mMap struct {
	mu sync.Mutex
	m  map[interface{}]int
}

var mm = &mMap{mu: sync.Mutex{}, m: map[interface{}]int{}}
var rwm = &rwMap{rw: sync.RWMutex{}, m: map[interface{}]int{}}
var syncm = &sync.Map{}

func ReadMM(key interface{}) {
	mm.mu.Lock()
	_ = mm.m[key]
	mm.mu.Unlock()
}

func WriteMM(key interface{}) {
	mm.mu.Lock()
	mm.m[key] = 10000
	mm.mu.Unlock()
}

func DeleteMM(key interface{}) {
	mm.mu.Lock()
	if _, ok := mm.m[key]; !ok {
		mm.mu.Unlock()
		return
	} else {
		delete(mm.m, key)
	}
	mm.mu.Unlock()
}

func ReadRWM(key interface{}) {
	rwm.rw.RLock()
	_ = rwm.m[key]
	rwm.rw.RUnlock()
}

func WriteRWM(key interface{}) {
	rwm.rw.Lock()
	rwm.m[key] = 10000
	rwm.rw.Unlock()
}

func DeleteRWM(key interface{}) {
	rwm.rw.Lock()
	if _, ok := rwm.m[key]; !ok {
		rwm.rw.Unlock()
		return
	} else {
		delete(rwm.m, key)
	}
	rwm.rw.Unlock()
}

func ReadSyncM(key interface{}) {
	_, _ = syncm.Load(key)
}

func WriteSyncM(key interface{}) {
	syncm.Store(key, 10000)
}

func DeleteSyncM(key interface{}) {
	syncm.Delete(key)
}
