package dawn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ContentType string

const (
	ContentTypeKey = "Content-Type"

	ContentString ContentType = "text/plain"
	ContentJSON   ContentType = "application/json"
	ContentHTML   ContentType = "text/html"
)

type H map[string]any

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) String(code int, format string, a ...any) {
	c.SetHeader(ContentTypeKey, string(ContentString))
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, a...)))
}

func (c *Context) JSON(code int, v any) {
	c.SetHeader(ContentTypeKey, string(ContentJSON))
	c.Status(code)
	if err := json.NewEncoder(c.Writer).Encode(v); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader(ContentTypeKey, string(ContentHTML))
	c.Status(code)
	c.Writer.Write([]byte(html))
}
