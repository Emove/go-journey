package context

import "testing"

func TestContextPropagation(t *testing.T) {
	ContextPropagation()
}

func TestStopParallelGoroutine(t *testing.T) {
	StopParallelGoroutine()
}

func TestCancelByChildrenContext(t *testing.T) {
	CancelByChildrenContext()
}
