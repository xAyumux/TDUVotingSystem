package service

import (
	"github.com/xAyumux/TDUVotingSystem/database"
	"github.com/xAyumux/TDUVotingSystem/model"
)

func GetPollService() []model.Polls {
	results := database.SelectPolls()
	return results
}

func GetQueryService(id int64) []model.Polls {
	results := database.SelectQuery(id)
	return results
}

func GetListService() []model.Votes {
	results := database.SelectVotes()
	return results
}

func PostPollsService(id int64, title, description string) int64 {
	lastInsertID := database.InsertPolls(id, title, description)
	return lastInsertID
}

func PostVotesService(userid string, id, result int64) int64 {
	lastInsertID := database.InsertVotes(userid, id, result)
	return lastInsertID
}
