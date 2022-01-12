package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xAyumux/TDUVotingSystem/service"
)

func GetPoll(c *gin.Context) {
	results, number := service.GetPollService()

	c.JSON(http.StatusOK, gin.H{
		"message": "GET Poll OK!",
		"results": results,
		"number":  number,
	})
}

func GetQuery(c *gin.Context) {
	idString := c.Query("id")
	var id int64
	id, _ = strconv.ParseInt(idString, 10, 64)
	results, number := service.GetQueryService(id)

	c.JSON(http.StatusOK, gin.H{
		"message":     "GET QueryString Poll OK!",
		"QueryString": id,
		"results":     results,
		"number":      number,
	})
}

func GetPollsList(c *gin.Context) {
	results := service.GetPollsListService()

	c.JSON(http.StatusOK, gin.H{
		"message": "GET List OK!",
		"results": results,
	})
}

func PostPolls(c *gin.Context) {
	idString := c.PostForm("id")
	var id int64
	id, _ = strconv.ParseInt(idString, 10, 64)
	title := c.Query("title")
	description := c.PostForm("description")

	lastInsertID := service.PostPollsService(id, title, description)

	c.JSON(http.StatusOK, gin.H{
		"message":             "POST Polls OK!",
		"formdataID":          id,
		"formdataTitle":       title,
		"formdataDescription": description,
		"lastInsertID":        lastInsertID,
	})
}

func PostVotes(c *gin.Context) {
	userid := c.PostForm("userid")
	idString := c.Query("id")
	var id int64
	id, _ = strconv.ParseInt(idString, 10, 64)
	resultString := c.PostForm("result")
	var result int64
	result, _ = strconv.ParseInt(resultString, 10, 64)

	lastInsertID := service.PostVotesService(userid, id, result)

	c.JSON(http.StatusOK, gin.H{
		"message":        "POST Votes OK!",
		"formdataUserID": userid,
		"formdataID":     id,
		"formdataResult": result,
		"lastInsertID":   lastInsertID,
	})
}
