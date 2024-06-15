package models

type Prediction struct {
	PredictionID int       `json:"prediction_id"`
	GoalsLocal   int       `json:"goals_local"`
	GoalsVisitor int       `json:"goals_visitor"`
	DocumentID   string    `json:"document_id"`
	MatchID      int       `json:"match_id"`
}
