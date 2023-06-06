package context

import (
	"context"
	"testing"
	"time"
)

func TestDeadlineContext(t *testing.T) {
	DeadlineContext()
}

func TestTimeoutContext(t *testing.T) {
	TimeoutContext()
}

func TestCancelBeforeDone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	CancelBeforeDone(ctx)
	select {
	case <-ctx.Done():
		cancel()
		time.Sleep(1 * time.Second)
		return
	}
}
