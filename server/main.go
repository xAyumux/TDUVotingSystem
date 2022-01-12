package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.Use(middleware.RecordUaAndTime)

	router.GET("/get/poll", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET Poll OK!",
		})
	})
	router.GET("/poll", func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"message":     "GET QueryString Poll OK!",
			"QueryString": id,
		})
	})
	router.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET List OK!",
		})
	})
	router.POST("/vote", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST Vote OK!",
		})
	})

	router.Run(":8000")
}
