package router

import (
	"container/list"
	"net/http"
	"sync"

	"github.com/s-beats/rest-todo/server"
)

type router struct {
	handlerRWMu  sync.RWMutex
	handlersGET  map[string]http.HandlerFunc
	handlersPOST map[string]http.HandlerFunc

	middlewareRWMu      sync.RWMutex
	middlewareFunctions *list.List
}

func NewRouter() *router {
	return newRouter()
}

func newRouter() *router {
	return &router{
		handlersGET:         make(map[string]http.HandlerFunc),
		handlersPOST:        make(map[string]http.HandlerFunc),
		middlewareFunctions: list.New(),
	}
}

func (r *router) handlers(method string) map[string]http.HandlerFunc {
	switch method {
	case http.MethodGet:
		return r.handlersGET
	case http.MethodPost:
		return r.handlersPOST
	default:
		panic("invalid method")
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.handlerRWMu.RLock()
	defer r.handlerRWMu.RUnlock()
	r.middlewareRWMu.RLock()
	defer r.middlewareRWMu.RUnlock()

	handlers := r.handlers(req.Method)
	fn, ok := handlers[req.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for e := r.middlewareFunctions.Back(); e != nil; e = e.Prev() {
		fn = e.Value.(server.Middleware)(fn)
	}

	fn(w, req)
}

// TODO: 重複の対応
func (r *router) PushBackMiddleware(m server.Middleware) server.Router {
	r.middlewareRWMu.Lock()
	defer r.middlewareRWMu.Unlock()

	if m == nil {
		return r
	}

	r.middlewareFunctions.PushBack(m)
	return r
}

func (r *router) GET(pattern string, fn http.HandlerFunc) {
	r.handlerRWMu.Lock()
	defer r.handlerRWMu.Unlock()

	r.handlersGET[pattern] = fn
}

func (r *router) POST(pattern string, fn http.HandlerFunc) {
	r.handlerRWMu.Lock()
	defer r.handlerRWMu.Unlock()

	r.handlersPOST[pattern] = fn
}
