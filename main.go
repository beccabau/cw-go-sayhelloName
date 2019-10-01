package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var httpAddr = flag.String("http.addr", ":8000", "HTTP listen address")

func main() {
	ge := gin.Default()

	// so default case works and app will run.
	ge.GET("/", func(c *gin.Context) {
		// c.Status(200)
		c.String(http.StatusOK, "Append /hello to the URL to get started.")
	})
	ge.GET("/hello", hello)
	ge.GET("/hello/:name", helloName)

	ge.Run(*httpAddr)
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello there! Append your name for a more personalized greeting.")
}

func helloName(c *gin.Context) {
	name := c.Param("name")
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05 Monday")
	c.String(http.StatusOK, "Hello %s. It is currently %s", name, formattedTime)
}
