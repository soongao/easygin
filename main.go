package main

import (
	"frame"
	"log"
	"net/http"
	"time"
)

func onlyForV2() frame.HandleFunc {
	return func(c *frame.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	// sort.Reverse()
	r := frame.Default()
	r.Use(frame.Logger()) // global midlleware
	r.GET("/", func(c *frame.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *frame.Context) {
			// expect /hello/me
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")

}
