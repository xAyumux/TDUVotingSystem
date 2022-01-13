package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xAyumux/TDUVotingSystem/service"
)

func GetPoll(c *gin.Context) {
	results, resultsVotes := service.GetPollService()

	c.JSON(http.StatusOK, gin.H{
		"message":           "GET Poll OK!",
		"results":           results,
		"resultId":          results.Id,
		"resultTitle":       results.Title,
		"resultDescription": results.Description,
		"resultsVotes":      resultsVotes,
	})
}

func GetQuery(c *gin.Context) {
	id := c.Query("id")
	results, resultsVotes := service.GetQueryService(id)

	c.JSON(http.StatusOK, gin.H{
		"message":           "GET QueryString Poll OK!",
		"QueryString":       id,
		"results":           results,
		"resultId":          results.Id,
		"resultTitle":       results.Title,
		"resultDescription": results.Description,
		"resultsVotes":      resultsVotes,
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
	id := c.PostForm("id")
	title := c.PostForm("title")
	description := c.PostForm("description")

	lastInsertID := service.PostPollsService(title, description)

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
	id := c.PostForm("id")
	result := c.PostForm("result")

	lastInsertID := service.PostVotesService(userid, id, result)

	c.JSON(http.StatusOK, gin.H{
		"message":        "POST Votes OK!",
		"formdataUserID": userid,
		"formdataID":     id,
		"formdataResult": result,
		"lastInsertID":   lastInsertID,
	})
}
