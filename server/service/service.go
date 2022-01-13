package service

import (
	"github.com/xAyumux/TDUVotingSystem/database"
	"github.com/xAyumux/TDUVotingSystem/model"
)

func GetPollService() (model.Polls, []model.Votes) {
	results, resultsVotes := database.SelectPoll()
	return results, resultsVotes
}

func GetQueryService(id string) (model.Polls, []model.Votes) {
	results, resultsVotes := database.SelectQuery(id)
	return results, resultsVotes
}

func GetPollsListService() []model.Polls {
	results := database.SelectPollsList()
	return results
}

func PostPollsService(title, description string) int64 {
	lastInsertID := database.InsertPolls(title, description)
	return lastInsertID
}

func PostVotesService(userid, id, result string) int64 {
	lastInsertID := database.InsertVotes(userid, id, result)
	return lastInsertID
}
