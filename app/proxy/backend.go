package proxy

import (
	"net/url"
	"sync"
)

type resolverBinding struct {
	resolve func(*url.URL) (*url.URL, error)
}

var externalResolver struct {
	sync.RWMutex
	current *resolverBinding
}

// UseResolver installs an external host's system proxy reader. Environment
// overrides retain priority. The reader must be safe for concurrent calls;
// releasing a binding stops new calls but does not cancel calls in progress.
// Release is idempotent and cannot remove a newer binding.
func UseResolver(resolve func(*url.URL) (*url.URL, error)) func() {
	b := &resolverBinding{resolve: resolve}
	externalResolver.Lock()
	externalResolver.current = b
	externalResolver.Unlock()
	var once sync.Once
	return func() {
		once.Do(func() {
			externalResolver.Lock()
			if externalResolver.current == b {
				externalResolver.current = nil
			}
			externalResolver.Unlock()
		})
	}
}

func selectedResolver() func(*url.URL) (*url.URL, error) {
	externalResolver.RLock()
	defer externalResolver.RUnlock()
	if b := externalResolver.current; b != nil {
		return b.resolve
	}
	return nil
}
