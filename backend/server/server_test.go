package server

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type testRouter struct{}

func (t *testRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// NOOP
	http.DefaultServeMux.ServeHTTP(w, req)
}
func (t *testRouter) GET(pattern string, fn http.HandlerFunc) {
	http.Handle(pattern, fn)
}
func (t *testRouter) POST(pattern string, fn http.HandlerFunc) {
	// NOOP
}
func (t *testRouter) PushBackMiddleware(m Middleware) Router {
	// NOOP
	return t
}

func TestStart(t *testing.T) {
	t.Setenv("HOST", "127.0.0.1")
	t.Setenv("PORT", "10000")
	r := &testRouter{}
	idleChanClosed := make(chan struct{})
	defer func() {
		<-idleChanClosed
	}()
	if err := Start(r); err != nil {
		t.Errorf("Start() error = %v", err)
	}
	resp, err := http.Get("127.0.0.1/10000/health")
	if err != nil {
		t.Errorf("Request to 127.0.0.1/10000/health = %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read response body = %v", err)
	}
	if !bytes.Equal(body, []byte("OK")) {
		t.Errorf("Expected response body = %v, actual = %v", "OK", body)
	}
}
