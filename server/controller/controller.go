package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xAyumux/TDUVotingSystem/service"
)

func GetPoll(c *gin.Context) {
	results := service.GetPollService()

	c.JSON(http.StatusOK, gin.H{
		"message": "GET Poll OK!",
		"results": results,
	})
}

func GetQuery(c *gin.Context) {
	idString := c.Query("id")
	var id int64
	id, _ = strconv.ParseInt(idString, 10, 64)
	results := service.GetQueryService(id)

	c.JSON(http.StatusOK, gin.H{
		"message":     "GET QueryString Poll OK!",
		"QueryString": id,
		"results":     results,
	})
}

func GetList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET List OK!",
	})
}

func PostVote(c *gin.Context) {
	userid := c.PostForm("userid")
	idString := c.Query("id")
	var id int64
	id, _ = strconv.ParseInt(idString, 10, 64)
	resultString := c.PostForm("result")
	var result int64
	result, _ = strconv.ParseInt(resultString, 10, 64)

	lastInsertID := service.PostVoteService(userid, id, result)

	c.JSON(http.StatusOK, gin.H{
		"message":        "POST Vote OK!",
		"formdataUserID": userid,
		"formdataID":     id,
		"formdataResult": result,
		"lastInsertID":   lastInsertID,
	})
}