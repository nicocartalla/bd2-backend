package models

type Championship struct {
    ChampionshipID   int    `json:"championship_id"`
    Name             string `json:"name"`
    Year             int    `json:"year"`
    Country          string `json:"country"`
    ChampionshipType string `json:"championship_type"`
    Champion         *int   `json:"champion"`
    Subchampion      *int   `json:"subchampion"`
}
