package context

import "testing"

func TestDeadlineContext(t *testing.T) {
	DeadlineContext()
}

func TestTimeoutContext(t *testing.T) {
	TimeoutContext()
}
