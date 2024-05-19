package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Init a new gin router and start basic api with a single endpoint
func main() {
	gin.ForceConsoleColor()

	router := gin.Default()

	router.POST("/api/v1/echo", func(c *gin.Context) {
		var json map[string]interface{}
		if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, json)
	})

	router.Run()
}
