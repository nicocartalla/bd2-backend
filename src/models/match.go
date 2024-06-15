package models
import (
	"time"
)



type Match struct {
	MatchID        	int 		`json:"match_id"`
	MatchDate      	time.Time	`json:"match_date"`
	TeamLocalID     int 		`json:"team_local_id"`
	TeamVisitorID   int 		`json:"team_visitor_id"`
	GoalsLocal	    *int    	`json:"goals_local"`
	GoalsVisitor    *int    	`json:"goals_visitor"`
	ChampionshipID  int  		`json:"championship_id"`
	StageID         *int  		`json:"stage_id"`
	GroupSID        *int  		`json:"group_s_id"`
}
