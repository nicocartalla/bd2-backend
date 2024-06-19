package models

import (
	"time"
)

type Notification struct {
	ID          int
	UserId 		int
	MatchId 	int
	NotificationTime time.Time
	NotificationMethod string
}

type UsersWhoHaveNotMadePredictionsTemplate struct {
	UserId 		 	 int
	MatchId 	 	 int
	UserEmail        string
	UserFirstName    string
	UserLastName     string
	ChampionshipName string
	LocalTeam        string
	VisitorTeam      string
	MatchDate        string
	LocalTeamLogo    string
	VisitorTeamLogo  string
}