package main

import (
	"dawn/dawn"
	"fmt"
)

const _testListenAddr = "192.168.204.130:5432"

func main() {
	panicTest()
}

func panicTest() {
	r := dawn.Default()
	r.GET("/panic", func(c *dawn.Context) {
		s := []int{1}
		fmt.Println(s[10])
		c.String(200, "tutu")
	})
	r.Run(_testListenAddr)
}

func staticTest() {
	r := dawn.New()
	r.Static("/assert", "/home/dawn/tmp")
	r.Run(_testListenAddr)
}
