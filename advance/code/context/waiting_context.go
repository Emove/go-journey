package context

import (
	"context"
	"time"
)

type WaitingContext struct {
	timeout    bool
	finishChan chan int
	val        interface{}
	ctx        context.Context
	cf         context.CancelFunc
}

func NewWaitingContext(duration time.Duration) *WaitingContext {
	ctx, cancelFunc := context.WithTimeout(context.Background(), duration)
	return &WaitingContext{
		timeout:    false,
		ctx:        ctx,
		cf:         cancelFunc,
		finishChan: make(chan int, 1),
	}
}

func (ctx *WaitingContext) SetValue(val interface{}) {
	ctx.val = val
	ctx.finishChan <- 1
}

func (ctx *WaitingContext) GetValue() interface{} {
	return ctx.val
}

func (ctx *WaitingContext) Cancel() {
	ctx.cf()
}

func (ctx *WaitingContext) IsTimeout() bool {
	return ctx.timeout
}

func (ctx *WaitingContext) Done() {
	defer ctx.cf()
	go func() {
		for {
			select {
			case <-ctx.ctx.Done():
				if ctx.ctx.Err() != nil {
					ctx.timeout = true
				}
				ctx.finishChan <- 1
				return
			}
		}
	}()
	for {
		select {
		case <-ctx.finishChan:
			return
		}
	}
}
