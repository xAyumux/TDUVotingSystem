package model

type Polls struct {
	Id          int64
	Title       string
	Description string
}

type Votes struct {
	UserID string
	Id     int64
	Result int64
}
