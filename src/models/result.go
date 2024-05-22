package models

type Result struct {
	ResultID         int    `json:"result_id"`
	MatchID       string `json:"match_id"`
	GoalsLocal	  int    `json:"goals_local"`
	GoalsVisitor int    `json:"goals_visitor"`
}
