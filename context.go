package dawn

import (
	"dawn/binding"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"sync"
)

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = binding.MIMEJSON
	MIMEHTML              = binding.MIMEJSON
	MIMEXML               = binding.MIMEJSON
	MIMEXML2              = binding.MIMEJSON
	MIMEPlain             = binding.MIMEJSON
	MIMEPOSTForm          = binding.MIMEJSON
	MIMEMultipartPOSTForm = binding.MIMEJSON
	MIMEPROTOBUF          = binding.MIMEJSON
	MIMEMSGPACK           = binding.MIMEJSON
	MIMEMSGPACK2          = binding.MIMEJSON
	MIMEYAML              = binding.MIMEJSON
	MIMETOML              = binding.MIMEJSON
)

const (
	// BodyBytesKey indicates a default body bytes key.
	BodyBytesKey = "_dawn/bodybyteskey"

	// ContextKey is the key that a Context returns itself for.
	ContextKey = "_dawn/contextkey"
)

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1

// Context allows us to pass variables between middleware, manage the flow, validate the JSON
// of a request and render a JSON response for example.
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    string
	fullPath string

	engine       *Engine
	params       *Params
	skippedNodes *[]skippedNode

	// this mutex protects Keys map.
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]any

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Errors errorMsgs

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string

	// queryCache caches the query result from c.Request.URL.Query().
	queryCache url.Values

	// formCache caches c.Request.PostForm, which contains the parsed form data from POST, PATCH,
	// or PUT body parameters.
	formCache url.Values

	// sameSite allows a server to define a cookie attribute making it impossible for the browser
	// to send this cookie along with cross-site requests.
	sameSite http.SameSite
}

type Contextv struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index    int

	engine *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

type H map[string]any

func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
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

func (c *Context) HTML(code int, name string, data any) {
	c.SetHeader(ContentTypeKey, string(ContentHTML))
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}
