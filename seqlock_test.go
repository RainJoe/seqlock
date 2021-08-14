package seqlock

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {
	l := NewSeqLock()
	wg := sync.WaitGroup{}
	sequence := make([]int, 0)
	n := 1000
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			l.Lock(uint64(i))
			sequence = append(sequence, i)
			l.Unlock(uint64(i))
		}(i, &wg)
	}
	wg.Wait()
	for idx, v := range sequence {
		if idx != v {
			t.Errorf("fail: expected: %v got: %v\n", idx, v)
		}
	}
}
