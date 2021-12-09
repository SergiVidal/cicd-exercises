package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func main() {
	r := gin.Default()
	r.GET("/demo", HelloWorld)
	r.Run()
}
