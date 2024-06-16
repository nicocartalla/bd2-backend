package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/utils"
	"fmt"
)

type ScoreChampionshipService struct{}


func (s *ScoreChampionshipService) InsertOrUpdateScoreChampionship(DocumentID string, championshipID int, points int) error {
	query := `
		INSERT INTO ScoresChampionship (document_id, championship_id, points) 
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE points = VALUES(points)
	`
	_, err := database.InsertDBParams(query, DocumentID, championshipID, points)
	if err != nil {
		utils.ErrorLogger.Println("Error inserting or updating score championship: ", err)
		return fmt.Errorf("error inserting or updating score championship: %v", err)
	}
	return nil
}

