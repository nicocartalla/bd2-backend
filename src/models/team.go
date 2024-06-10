package models

type Team struct {
	ID         int    `json:"team_id"`
	Name       string `json:"name"`
	URLLogo    string `json:"url_logo"`
	Description string `json:"description"`
}
