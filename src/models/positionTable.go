package models

type PositionTable struct {
	DocumentID      string `json:"document_id"`
	FirstName		string `json:"first_name"`
	LastName		string `json:"last_name"`
	ChampionshipID  int    `json:"championship_id"`
	Points	   	  	int    `json:"points"`
}
