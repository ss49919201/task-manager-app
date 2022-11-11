package ginrouter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s-beats/rest-todo/server"
)

// TODO
// gin はミドルウェア->ハンドラーの順番で準備しないといけないので、他と同じように使えない

type router struct {
	ginEngine *gin.Engine
}

func NewRouter() *router {
	return newRouter()
}

func newRouter() *router {
	return &router{
		ginEngine: gin.Default(),
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.ginEngine.ServeHTTP(w, req)
}

func (r *router) PushBackMiddleware(m server.Middleware) server.Router {
	r.ginEngine.Use(func(c *gin.Context) {
		m(func(w http.ResponseWriter, req *http.Request) {
			c.Next()
		})(c.Writer, c.Request)
	})
	l := len(r.ginEngine.Handlers)
	fmt.Println(l)
	return r
}

func (r *router) Get(pattern string, fn http.HandlerFunc) {
	r.ginEngine.GET(pattern, func(c *gin.Context) {
		fn(c.Writer, c.Request)
	})
}

func (r *router) Post(pattern string, fn http.HandlerFunc) {
	r.ginEngine.POST(pattern, func(c *gin.Context) {
		fn(c.Writer, c.Request)
	})
}
