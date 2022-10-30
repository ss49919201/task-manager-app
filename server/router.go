package server

import (
	"net/http"
	"sync"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

type router struct {
	handlersGET         map[string]http.HandlerFunc
	handlersPOST        map[string]http.HandlerFunc
	middlewareFunctions []middleware

	rwMu sync.RWMutex
}

func NewRouter() *router {
	return &router{
		handlersGET:         make(map[string]http.HandlerFunc),
		handlersPOST:        make(map[string]http.HandlerFunc),
		middlewareFunctions: []middleware{},
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

	for _, m := range r.middlewareFunctions {
		fn = m(fn)
	}

	fn(w, req)
}

// TODO: 重複の対応
func (r *router) SetMiddleware(m middleware) *router {
	r.rwMu.Lock()
	defer r.rwMu.Unlock()

	if m == nil {
		return r
	}

	r.middlewareFunctions = append(r.middlewareFunctions, m)
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
