package server

import "net/http"

type middleware func(http.HandlerFunc) http.HandlerFunc

type router struct {
	middlewareFunctions []middleware
}

func NewRouter() *router {
	return &router{}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO:
}

// TODO: 重複の対応
func (r *router) SetMiddleware(m middleware) *router {
	if m == nil {
		return r
	}

	r.middlewareFunctions = append(r.middlewareFunctions, m)
	return r
}

func (r *router) Get(pattern string, fn http.HandlerFunc) {
	// TODO:
}
func (r *router) Post(pattern string, fn http.HandlerFunc) {
	// TODO:
}
