package proxy

import (
	"net/url"
	"testing"
)

func TestExternalResolverReleaseCannotRemoveReplacement(t *testing.T) {
	for _, key := range []string{"HTTP_PROXY", "http_proxy", "HTTPS_PROXY", "https_proxy", "ALL_PROXY", "all_proxy"} {
		t.Setenv(key, "")
	}
	target, _ := url.Parse("https://example.test")
	first := UseResolver(func(*url.URL) (*url.URL, error) { t.Fatal("old resolver used"); return nil, nil })
	called := 0
	second := UseResolver(func(got *url.URL) (*url.URL, error) {
		called++
		if got != target {
			t.Fatal("target replaced")
		}
		return &url.URL{Scheme: "http", Host: "proxy.test:8080"}, nil
	})
	t.Cleanup(second)
	first()
	first()
	got, err := SystemProxy(target)
	if err != nil || got == nil || got.Host != "proxy.test:8080" || called != 1 {
		t.Fatalf("%v %v calls=%d", got, err, called)
	}
	second()
	second()
	if selectedResolver() != nil {
		t.Fatal("released resolver retained")
	}
}
