package models

type PredictionChampionship struct {
	PredictionCID 	int       `json:"prediction_c_id"`
	Champion   		int       `json:"champion"`
	SubChampion 	int       `json:"subchampion"`
	DocumentID   	string    `json:"document_id"`
	ChampionshipID 	int       `json:"championship_id"`
}
