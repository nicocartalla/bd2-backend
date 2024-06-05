package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
)

type PositionTableService struct{}


func (s *PositionTableService) GetPositionTableByChampionship(championshipID int) ([]models.PositionTable, error) {
	query := `
		SELECT u.document_id, u.first_name, u.last_name, gm.championship_id, SUM(s.points) AS points
		FROM User u
		JOIN Scores s ON u.document_id = s.document_id
		JOIN GameMatch gm ON s.match_id = gm.match_id
		WHERE gm.championship_id = ?
		GROUP BY u.document_id, gm.championship_id
		ORDER BY points DESC
	`
	rows, err := database.QueryRowsDBParams(query, championshipID)
	if err != nil {
		utils.ErrorLogger.Println("Error querying leaderboard by championship:", err)
		return nil, fmt.Errorf("error querying leaderboard by championship: %v", err)
	}
	defer rows.Close()

	var leaderboard []models.PositionTable
	for rows.Next() {
		var userScore models.PositionTable
		if err := rows.Scan(&userScore.DocumentID, &userScore.FirstName, &userScore.LastName, &userScore.ChampionshipID, &userScore.Points); err != nil {
			utils.ErrorLogger.Println("Error scanning user score:", err)
			return nil, fmt.Errorf("error scanning user score: %v", err)
		}
		leaderboard = append(leaderboard, userScore)
	}
	return leaderboard, nil
}
