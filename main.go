package main

import (
	"dawn/dawn"
	"fmt"
	"net/http"
)

const _testListenAddr = "192.168.204.130:5432"

func pathHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.URL.Path)
}

func hellpHandleFunc(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintln(w, k, v)
	}
}

func main() {
	engine := dawn.New()
	engine.Get("/", pathHandleFunc)
	engine.Get("/hello", hellpHandleFunc)

	engine.Run(_testListenAddr)
}
