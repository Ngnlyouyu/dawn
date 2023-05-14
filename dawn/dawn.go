package dawn

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandleFunc
}

var _ http.Handler = &Engine{}

// New is a constructor of Engine
func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

func (e *Engine) addRoute(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	fmt.Println("add route: ", key)
	e.router[key] = handler
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
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintln(w, "404 not found")
	}
}
