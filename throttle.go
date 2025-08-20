package generic

import (
	"sync/atomic"
	"time"
)

type Throttler struct {
	scheduled atomic.Bool
	fn        func()
	delay     time.Duration
}

func NewThrottler(delay time.Duration, fn func()) *Throttler {
	return &Throttler{
		fn:    fn,
		delay: delay,
	}
}

func (t *Throttler) ThrottledCall() {
	if t.scheduled.CompareAndSwap(false, true) {
		time.AfterFunc(t.delay, func() {
			t.scheduled.Store(false)
			t.fn()
		})
	}
}
