# seqlock


In some special situations, we need to make task execution in order through goroutines. For example, if we need to dispatch tasks to goroutines and some of the tasks need to be executed sequentially，in this situation, we can give these tasks a seqlock to ensure order，and other tasks can execution concurrently.


## usage

```golang
l := NewSeqLock()
// seq is the task order, it's should start from zero and be monotonically increasing
l.Lock(seq)
// your task
l.Unlock(seq)
```
