package router

import (
	"bytes"
	"container/list"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_router(t *testing.T) {
	r := NewRouter()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
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
	r.GET("/ping", func(w http.ResponseWriter, r *http.Request) {
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
	r.POST("/test", func(w http.ResponseWriter, r *http.Request) {
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

func Test_router_middlware(t *testing.T) {
	r := NewRouter()
	// 1,2,3
	r.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/" + strconv.Itoa(2)))
	})
	r.
		PushBackMiddleware(func(next http.HandlerFunc) http.HandlerFunc {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(strconv.Itoa(1)))
				next.ServeHTTP(w, r)
			})
		}).
		PushBackMiddleware(func(next http.HandlerFunc) http.HandlerFunc {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
				w.Write([]byte("/" + strconv.Itoa(3)))
			})
		})

	testServer := httptest.NewServer(r)
	defer testServer.Close()
	// GET
	resp, err := http.Get(testServer.URL + "/test")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	timestampBody := strings.Split(string(respBody), "/")
	one, err := strconv.Atoi(timestampBody[0])
	assert.NoError(t, err)
	two, err := strconv.Atoi(timestampBody[1])
	assert.NoError(t, err)
	three, err := strconv.Atoi(timestampBody[2])
	assert.NoError(t, err)
	assert.True(t, three > two && two > one)
}

func Test_router_PushBackMiddleware(t *testing.T) {
	testFn := func(f http.HandlerFunc) http.HandlerFunc { return f }
	hasTestFnRouter := newRouter()
	hasTestFnRouter.middlewareFunctions.PushBack(testFn)

	type fields struct {
		middlewareFunctions *list.List
	}
	type args struct {
		m func(f http.HandlerFunc) http.HandlerFunc
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
			hasTestFnRouter,
		},
		{
			"add nil",
			fields{},
			args{nil},
			&router{
				middlewareFunctions: list.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				middlewareFunctions: tt.fields.middlewareFunctions,
			}
			_ = r.PushBackMiddleware(tt.args.m)
			assert.Equal(t, tt.want.middlewareFunctions.Len(), r.middlewareFunctions.Len())
		})
	}
}
