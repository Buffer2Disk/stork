package mtx

import (
	"github.com/argcv/stork/log"
	"runtime"
	"sync/atomic"
)

// Compare with waiting group
// it will return current state
// aka **How many workers are still working**
type WaitGroupWithState struct {
	st int64
}

func NewWaitGroupWithState() *WaitGroupWithState {
	return &WaitGroupWithState{
		st: 0,
	}
}

func (wg *WaitGroupWithState) Add(delta int64) int64 {
	newSt := atomic.AddInt64(&(wg.st), delta)
	if newSt < 0 {
		log.Fatalf("ERROR: status is lower than 0!!! (%v)", newSt)
	}
	return newSt
}

// minus one, return current value
func (wg *WaitGroupWithState) Done() int64 {
	return wg.Add(-1)
}

func (wg *WaitGroupWithState) State() int64 {
	return atomic.LoadInt64(&(wg.st))
}

func (wg *WaitGroupWithState) Wait() {
	for wg.State() > 0 {
		runtime.Gosched()
	}
}
