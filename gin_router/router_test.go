package ginrouter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_router_middlware(t *testing.T) {
	r := NewRouter()
	// 1,2,3
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
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/" + strconv.Itoa(2)))
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
