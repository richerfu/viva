package proxy

import (
	"net/http"
	"net/url"
	"sync"
)

type transportBinding struct {
	wrap func(*http.Transport) http.RoundTripper
}

var externalTransport struct {
	sync.RWMutex
	current *transportBinding
}

// UseTransport installs an external host's transport wrapper. The host can
// preserve ordered proxy failover that a single URL resolver cannot express.
// The wrapper must not mutate base, and must be safe for concurrent requests.
// Existing transports keep their binding. The returned function removes only
// this registration and may be called repeatedly.
func UseTransport(wrap func(*http.Transport) http.RoundTripper) func() {
	b := &transportBinding{wrap}
	externalTransport.Lock()
	externalTransport.current = b
	externalTransport.Unlock()
	var once sync.Once
	return func() {
		once.Do(func() {
			externalTransport.Lock()
			defer externalTransport.Unlock()
			if externalTransport.current == b {
				externalTransport.current = nil
			}
		})
	}
}

// WrapTransport applies the registered host wrapper to base, or returns base
// unchanged. Configure base before wrapping it. Close idle connections through
// http.Client.CloseIdleConnections. A nil base uses the default system transport.
func WrapTransport(base *http.Transport) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport.(*http.Transport).Clone()
		base.Proxy = func(r *http.Request) (*url.URL, error) { return SystemProxy(r.URL) }
	}
	externalTransport.RLock()
	binding := externalTransport.current
	externalTransport.RUnlock()
	if binding != nil && binding.wrap != nil {
		return binding.wrap(base)
	}
	return base
}
