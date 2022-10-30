package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_router(t *testing.T) {
	r := NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Hello", "Hello, World!")
		w.Write([]byte("Hello, World!"))
		w.WriteHeader(http.StatusOK)
	})

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	resp, err := http.Get(testServer.URL)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Hello, World!", resp.Header.Get("X-Hello"))
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "Hello, World!", string(respBody))
}

func Test_router_GET(t *testing.T) {
	r := NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	// GET
	resp, err := http.Get(testServer.URL + "/ping")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "pong", string(respBody))

	resp, err = http.Get(testServer.URL + "/not_found")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	respBody, err = io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "", string(respBody))

	// POST
	resp, err = http.Post(testServer.URL, "application/json", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	respBody, err = io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "", string(respBody))
}

func Test_router_POST(t *testing.T) {
	r := NewRouter()
	r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		w.Write([]byte("success: " + string(b)))
		w.WriteHeader(http.StatusOK)
	})

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	// GET
	resp, err := http.Get(testServer.URL + "/ping")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "", string(respBody))

	// POST
	resp, err = http.Post(testServer.URL+"/test", "application/json", bytes.NewBuffer([]byte("test")))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	respBody, err = io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "success: test", string(respBody))

	resp, err = http.Post(testServer.URL+"/not_found", "application/json", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	respBody, err = io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "", string(respBody))
}

func Test_router_SetMiddleware(t *testing.T) {
	testFn := middleware(func(f http.HandlerFunc) http.HandlerFunc { return f })

	type fields struct {
		middlewareFunctions []middleware
	}
	type args struct {
		m middleware
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *router
	}{
		{
			"add one",
			fields{},
			args{testFn},
			&router{middlewareFunctions: []middleware{testFn}},
		},
		{
			"add nil",
			fields{},
			args{nil},
			&router{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				middlewareFunctions: tt.fields.middlewareFunctions,
			}
			got := r.SetMiddleware(tt.args.m)
			assert.Equal(t, len(tt.want.middlewareFunctions), len(got.middlewareFunctions))
		})
	}
}
