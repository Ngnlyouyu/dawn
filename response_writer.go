package dawn

import (
	"bufio"
	"net"
	"net/http"
)

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier

	// Status returns the HTTP response status code of the current request.
	Status() int

	// Size returns the number of bytes already written into the response http body.
	Size() int

	// WriteString writes the string into the response body.
	WriteString(string) (int, error)

	// Written returns true if the response body was already written.
	Written() bool

	// WriteHeaderNow forces to write the http header (status code + headers).
	WriteHeaderNow()

	// Pusher get the http.Pusher for server push.
	Pusher() http.Pusher
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

var _ ResponseWriter = (*responseWriter)(nil)

func (w *responseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func (w *responseWriter) reset(writer http.ResponseWriter) {}

func (w *responseWriter) WriteHeader(code int) {}

func (w *responseWriter) WriteHeaderNow() {}

func (w *responseWriter) Write(data []byte) (n int, err error) {
	return 0, nil
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
	return 0, nil
}

func (w *responseWriter) Status() int {
	return 0
}

func (w *responseWriter) Size() int {
	return 0
}

func (w *responseWriter) Written() bool {
	return false
}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

func (w *responseWriter) CloseNotify() <-chan bool {
	return nil
}

func (w *responseWriter) Flush() {}

func (w *responseWriter) Pusher() (pusher http.Pusher) {
	return nil
}
