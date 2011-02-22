package gc

import (
	"doozer/store"
	"github.com/bmizerany/assert"
	"testing"
)

// Testing

type FakeProposer chan string

func (fs FakeProposer) Propose(v []byte) (e store.Event) {
	fs <- string(v)
	e.Cas = 123
	return
}

func TestGcPulse(t *testing.T) {
	seqns := make(chan int64)
	defer close(seqns)
	fs := make(FakeProposer)

	go Pulse("test", seqns, fs, 1)

	seqns <- 0
	assert.Equal(t, "0:/doozer/info/test/applied=0", <-fs)

	seqns <- 1
	assert.Equal(t, "123:/doozer/info/test/applied=1", <-fs)
}
