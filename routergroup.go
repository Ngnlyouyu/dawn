package dawn

import (
	"net/http"
	"regexp"
)

var (
	// regEnLetter matches english letters for http method name
	regEnLetter = regexp.MustCompile("^[A-Z]+$")

	// anyMethod for RouterGroup Any Method
	anyMethod = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodConnect,
		http.MethodTrace,
	}
)

// IRouter defines all router handle interface includes single and group router.
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}

// IRoutes defines all router handle interface.
type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
	Match([]string, string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}

// RouterGroup is used internally to configure router, a RouterGroup is associated with
// a prefix and an array of handlers (middleware).
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}

var _ IRouter = (*RouterGroup)(nil)

func (g *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return nil
}

func (g *RouterGroup) BasePath() string {
	return ""
}

func (g *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	return nil
}

func (g *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) StaticFile(relativePath, filePath string) IRoutes {
	return nil
}

func (g *RouterGroup) StaticFileFS(relativePath, filePath string, fs http.FileSystem) IRoutes {
	return nil
}

func (g *RouterGroup) staticFileHandler(relativePath string, handler HandlerFunc) IRoutes {
	return nil
}

func (g *RouterGroup) Static(relativePath, root string) IRoutes {
	return nil
}

func (g *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	return nil
}

func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	return nil
}

func (g *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	return nil
}

func (g *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return ""
}

func (g *RouterGroup) returnObj() IRoutes {
	return nil
}
