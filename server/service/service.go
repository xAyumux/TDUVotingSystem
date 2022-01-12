package service

import (
	"github.com/xAyumux/TDUVotingSystem/database"
	"github.com/xAyumux/TDUVotingSystem/model"
)

func GetPollService() []model.Polls {
	results := database.Select()
	return results
}

func GetQueryService(id int64) []model.Polls {
	results := database.SelectQuery(id)
	return results
}

func GetListService() {
}

func PostVoteService(userid string, id, result int64) int64 {
	lastInsertID := database.Insert(userid, id, result)
	return lastInsertID
}
