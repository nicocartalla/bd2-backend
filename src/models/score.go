package models

type Score struct {
	ScoreID int    `json:"score_id"`
	UserID  string `json:"document_id"`
	MatchID string `json:"match_id"`
	Points  int    `json:"points"`
}
