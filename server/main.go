package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xAyumux/TDUVotingSystem/controller"
)

func main() {
	router := gin.Default()

	// idが一番大きいものを表示
	router.GET("/get/poll", controller.GetPoll)
	// 指定されたidを表示
	router.GET("/poll", controller.GetQuery)
	// title一覧を表示
	router.GET("/list", controller.GetList)
	// pollsテーブルに新しいデータを追加
	router.POST("/post/polls", controller.PostPolls)
	// votesテーブルに新しいデータを追加
	router.POST("/post/votes", controller.PostVotes)

	router.Run(":8000")
}
