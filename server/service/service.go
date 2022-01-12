package service

import (
	"github.com/xAyumux/TDUVotingSystem/database"
	"github.com/xAyumux/TDUVotingSystem/model"
)

func GetPollService() ([]model.Polls, int64) {
	results, number := database.SelectPoll()
	return results, number
}

func GetQueryService(id int64) ([]model.Polls, int64) {
	results, number := database.SelectQuery(id)
	return results, number
}

func GetPollsListService() []model.Polls {
	results := database.SelectPollsList()
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
