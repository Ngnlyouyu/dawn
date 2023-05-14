package main

import (
	"dawn/dawn"
	"net/http"
)

const _testListenAddr = "192.168.204.130:5432"

func pathHandleFunc(c *dawn.Context) {
	c.HTML(http.StatusOK, "<h1>Hello dawn</h1r>")
}

func helloHandleFunc(c *dawn.Context) {
	c.String(http.StatusOK, "path: %s, name: %s", c.Path, c.Query("name"))
}

func loginHandleFunc(c *dawn.Context) {
	c.JSON(http.StatusOK, dawn.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}

func helloNameHandleFunc(c *dawn.Context) {
	c.String(http.StatusOK, "param: %s", c.Param("name"))
}

func assetsHandleFunc(c *dawn.Context) {
	c.String(http.StatusOK, "filepath: %s", c.Param("filepath"))
}

func main() {
	engine := dawn.New()
	engine.Get("/", pathHandleFunc)
	engine.Get("/hello", helloHandleFunc)
	engine.Post("/login", loginHandleFunc)
	engine.Get("/hello/:name", helloNameHandleFunc)
	engine.Get("/assets/*filepath", assetsHandleFunc)

	engine.Run(_testListenAddr)
}
