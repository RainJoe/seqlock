package seqlock

import "sync"

type SeqLock struct {
	lock sync.Mutex
	cond *sync.Cond
	seq  uint64
}

func NewSeqLock() *SeqLock {
	l := &SeqLock{}
	l.cond = sync.NewCond(&l.lock)
	return l
}

func (s *SeqLock) Lock(i uint64) {
	s.lock.Lock()
	for i != s.seq {
		s.cond.Wait()
	}
	s.lock.Unlock()
}

func (s *SeqLock) Unlock(i uint64) {
	s.lock.Lock()
	s.seq++
	s.lock.Unlock()
	s.cond.Broadcast()
}
