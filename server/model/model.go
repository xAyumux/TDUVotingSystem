package model

type Polls struct {
	Id          int
	Title       string
	Description string
}

type Votes struct {
	UserID string
	Id     int
	Result int
}
