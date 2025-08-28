package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// setupRouter configures and returns a Gin engine.
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(); err != nil { // default :8080
		log.Fatalf("server failed: %v", err)
	}
}
