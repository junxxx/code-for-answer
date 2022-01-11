package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get is concurrency-safe! But duplicate suppression
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key
		memo.mu.Unlock()
		<-e.ready
	}

	return e.res.value, e.res.err
}
