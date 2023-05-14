package dawn

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandleFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

var _ http.Handler = &Engine{}

// New is a constructor of Engine
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) addRoute(method, pattern string, handler HandleFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) Get(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
