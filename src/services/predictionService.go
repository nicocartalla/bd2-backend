package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
)

type PredictionService struct {
	Prediction models.Prediction
}

func (s *PredictionService) GetPredictionsByUser(userID int) ([]models.Prediction, error) {
	query := `SELECT prediction_id, document_id, match_id, goals_local, goals_visitor FROM Predictions WHERE document_id = ?`
	rows, err := database.QueryRowsDBParams(query, userID)
	if err != nil {
		utils.ErrorLogger.Println("Error querying predictions by user:", err)
		return nil, fmt.Errorf("error querying predictions by user: %v", err)
	}
	defer rows.Close()

	var predictions []models.Prediction
	for rows.Next() {
		var prediction models.Prediction
		if err := rows.Scan(&prediction.PredictionID, &prediction.DocumentID, &prediction.MatchID, &prediction.GoalsLocal, &prediction.GoalsVisitor); err != nil {
			utils.ErrorLogger.Println("Error scanning prediction:", err)
			return nil, fmt.Errorf("error scanning prediction: %v", err)
		}
		predictions = append(predictions, prediction)
	}
	return predictions, nil
}


func (s *PredictionService) InsertPrediction(prediction models.Prediction) (int64, error) {
	query := `
		INSERT INTO Predictions (goals_local, goals_visitor, document_id, match_id, group_id)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		goals_local = VALUES(goals_local),
		goals_visitor = VALUES(goals_visitor),
		group_id = VALUES(group_id)`
	
	result, err := database.InsertDBParams(query, prediction.GoalsLocal, prediction.GoalsVisitor, prediction.DocumentID, prediction.MatchID, prediction.GroupID)
	if err != nil {
		utils.ErrorLogger.Println("Error inserting or updating prediction:", err)
		return 0, fmt.Errorf("error inserting or updating prediction: %v", err)
	}
	return result, nil
}

func (s *PredictionService) GetPredictionsByMatchID(matchID int) ([]models.Prediction, error) {
	query := "SELECT * FROM Predictions WHERE match_id = ?"
	rows, err := database.QueryRowsDBParams(query, matchID)
	if err != nil {
		utils.ErrorLogger.Println("Error getting predictions: ", err)
		return nil, fmt.Errorf("error getting predictions: %v", err)
	}
	defer rows.Close()

	var predictions []models.Prediction
	for rows.Next() {
		var prediction models.Prediction
		err = rows.Scan(&prediction.PredictionID, &prediction.GoalsLocal, &prediction.GoalsVisitor, &prediction.DocumentID, &prediction.MatchID, &prediction.GroupID)
		if err != nil {
			utils.ErrorLogger.Println("Error scanning prediction: ", err)
			return nil, fmt.Errorf("error scanning prediction: %v", err)
		}
		predictions = append(predictions, prediction)
	}
	return predictions, nil
}
