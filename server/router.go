package server

import (
	"container/list"
	"net/http"
	"sync"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

type router struct {
	handlersGET         map[string]http.HandlerFunc
	handlersPOST        map[string]http.HandlerFunc
	middlewareFunctions *list.List

	rwMu sync.RWMutex
}

func NewRouter() *router {
	return &router{
		handlersGET:         make(map[string]http.HandlerFunc),
		handlersPOST:        make(map[string]http.HandlerFunc),
		middlewareFunctions: list.New(),
	}
}

func (r *router) lazyInitMiddlewareFunctions() {
	if r.middlewareFunctions == nil {
		r.middlewareFunctions = list.New()
	}
}

func (r *router) handlers(method string) map[string]http.HandlerFunc {
	switch method {
	case http.MethodGet:
		return r.handlersGET
	case http.MethodPost:
		return r.handlersPOST
	default:
		return nil
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.rwMu.RLock()
	defer r.rwMu.RUnlock()

	handlers := r.handlers(req.Method)
	fn, ok := handlers[req.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for e := r.middlewareFunctions.Back(); e != nil; e = e.Prev() {
		fn = e.Value.(middleware)(fn)
	}

	fn(w, req)
}

// TODO: 重複の対応
func (r *router) PushBackMiddleware(m middleware) *router {
	r.lazyInitMiddlewareFunctions()

	r.rwMu.Lock()
	defer r.rwMu.Unlock()

	if m == nil {
		return r
	}

	r.middlewareFunctions.PushBack(m)
	return r
}

func (r *router) Get(pattern string, fn http.HandlerFunc) {
	r.rwMu.Lock()
	defer r.rwMu.Unlock()

	r.handlersGET[pattern] = fn
}

func (r *router) Post(pattern string, fn http.HandlerFunc) {
	r.rwMu.Lock()
	defer r.rwMu.Unlock()

	r.handlersPOST[pattern] = fn
}
