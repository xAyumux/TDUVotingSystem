package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xAyumux/TDUVotingSystem/controller"
)

func main() {
	router := gin.Default()

	// router.Use(middleware.RecordUaAndTime)

	router.GET("/get/poll", controller.GetPoll)
	router.GET("/poll", controller.GetQuery)
	router.GET("/list", controller.GetList)
	router.POST("/post/polls", controller.PostPolls)
	router.POST("/post/votes", controller.PostVotes)

	router.Run(":8000")
}
