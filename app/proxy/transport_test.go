package proxy

import (
	"net/http"
	"testing"
)

func TestExternalTransportOwnership(t *testing.T) {
	base := http.DefaultTransport.(*http.Transport).Clone()
	first := UseTransport(func(*http.Transport) http.RoundTripper { t.Fatal("replaced wrapper"); return nil })
	wrapped := base.Clone()
	second := UseTransport(func(got *http.Transport) http.RoundTripper {
		if got != base {
			t.Fatal("lost base")
		}
		return wrapped
	})
	defer second()
	first()
	first()
	if got := WrapTransport(base); got != wrapped {
		t.Fatal("wrapper not selected")
	}
	second()
	second()
	if got := WrapTransport(base); got != base {
		t.Fatal("wrapper retained")
	}
}
