package models

type Prediction struct {
	PredictionID int       `json:"prediction_id"`
	GoalsLocal   int       `json:"goals_local"`
	GoalsVisitor int       `json:"goals_visitor"`
	UserID       int       `json:"user_id"`
	MatchID      int       `json:"match_id"`
	GroupID      int       `json:"group_id"`
}
