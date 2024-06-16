package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
)

type PredictionChampionshipService struct {
	Prediction models.PredictionChampionship
}

func (s *PredictionChampionshipService) GetPredictionChampionshipByUser(documentID string, championshipID int) ([]models.PredictionChampionship, error) {
	query := `SELECT prediction_c_id, document_id, championship_id, champion, subchampion 
              FROM PredictionsChampionships 
              WHERE document_id = ? AND championship_id = ?`
	rows, err := database.QueryRowsDBParams(query, documentID, championshipID)
	if err != nil {
		utils.ErrorLogger.Println("Error querying predictions by user and championship:", err)
		return nil, fmt.Errorf("error querying predictions by user and championship: %v", err)
	}
	defer rows.Close()

	var predictions []models.PredictionChampionship
	for rows.Next() {
		var prediction models.PredictionChampionship
		if err := rows.Scan(&prediction.PredictionCID, &prediction.DocumentID, &prediction.ChampionshipID, &prediction.Champion, &prediction.SubChampion); err != nil {
			utils.ErrorLogger.Println("Error scanning prediction championship:", err)
			return nil, fmt.Errorf("error scanning prediction championship: %v", err)
		}
		predictions = append(predictions, prediction)
	}
	return predictions, nil
}


func (s *PredictionChampionshipService) InsertPredictionChampionship(prediction models.PredictionChampionship) (int64, error) {
	query := `
		INSERT INTO PredictionsChampionships (champion, subchampion, document_id, championship_id)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		champion = VALUES(champion),
		subchampion = VALUES(subchampion)`
	
	result, err := database.InsertDBParams(query, prediction.Champion, prediction.SubChampion, prediction.DocumentID, prediction.ChampionshipID)
	if err != nil {
		utils.ErrorLogger.Println("Error inserting or updating championship prediction:", err)
		return 0, fmt.Errorf("error inserting or updating championship prediction: %v", err)
	}
	return result, nil
}

func (s *PredictionChampionshipService) GetPredictionsChampionshipByChampionshipID(championshipID int) ([]models.PredictionChampionship, error) {
	query := `SELECT prediction_c_id, document_id, championship_id, champion, subchampion 
			  FROM PredictionsChampionships 
			  WHERE championship_id = ?`
	rows, err := database.QueryRowsDBParams(query, championshipID)
	if err != nil {
		utils.ErrorLogger.Println("Error querying predictions by championship:", err)
		return nil, fmt.Errorf("error querying predictions by championship: %v", err)
	}
	defer rows.Close()

	var predictions []models.PredictionChampionship
	for rows.Next() {
		var prediction models.PredictionChampionship
		if err := rows.Scan(&prediction.PredictionCID, &prediction.DocumentID, &prediction.ChampionshipID, &prediction.Champion, &prediction.SubChampion); err != nil {
			utils.ErrorLogger.Println("Error scanning prediction championship:", err)
			return nil, fmt.Errorf("error scanning prediction championship: %v", err)
		}
		predictions = append(predictions, prediction)
	}
	return predictions, nil
}