package model

type Score struct {
	ScoreID int    `json:"score_id"`
	UserID  string `json:"user_id"`
	Point   int    `json:"point"`
	Level   string `json:"level"`
}

type RequestScore struct {
	Score
	GuestName string `json:"guest_name"`
}
