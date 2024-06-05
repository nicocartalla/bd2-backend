package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/utils"
	"fmt"
)

type ScoreService struct{}

func (s *ScoreService) InsertOrUpdateScore(DocumentID string, matchID int, points int) error {
	query := `
		INSERT INTO Scores (document_id, match_id, points) 
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE points = VALUES(points)
	`
	_, err := database.InsertDBParams(query, DocumentID, matchID, points)
	if err != nil {
		utils.ErrorLogger.Println("Error inserting or updating score: ", err)
		return fmt.Errorf("error inserting or updating score: %v", err)
	}
	return nil
}

