package models

type Utils struct {
	ExactMatchPoints         int `json:"exact_match_points"`
	CorrectResultMatchPoints int `json:"correct_result_match_points"`
	ChampionPoints           int `json:"champion_points"`
	SubChampionPoints        int `json:"sub_champion_points"`
}
