package sync

import "sync"

func LockAfterRLock() {
	mu := sync.RWMutex{}
	mu.RLock()
	mu.Lock()

	defer mu.Unlock()
	defer mu.RUnlock()
}
