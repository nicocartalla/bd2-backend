package services

import (
	"bd2-backend/src/database"
	//"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
)

type PositionTableService struct{}

type UserScore struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Points    int    `json:"points"`
}

type GroupUserScore struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	GroupID   int    `json:"group_id"`
	Points    int    `json:"points"`
}

func (s *PositionTableService) GetPositionTableByChampionship(championshipID int) ([]UserScore, error) {
	query := `
		SELECT u.user_id, u.first_name, u.last_name, SUM(s.points) AS points
		FROM User u
		JOIN Scores s ON u.user_id = s.user_id
		JOIN GameMatch gm ON s.match_id = gm.match_id
		WHERE gm.championship_id = ?
		GROUP BY u.user_id, u.first_name, u.last_name
		ORDER BY points DESC
	`
	rows, err := database.QueryRowsDBParams(query, championshipID)
	if err != nil {
		utils.ErrorLogger.Println("Error querying leaderboard by championship:", err)
		return nil, fmt.Errorf("error querying leaderboard by championship: %v", err)
	}
	defer rows.Close()

	var leaderboard []UserScore
	for rows.Next() {
		var userScore UserScore
		if err := rows.Scan(&userScore.UserID, &userScore.FirstName, &userScore.LastName, &userScore.Points); err != nil {
			utils.ErrorLogger.Println("Error scanning user score:", err)
			return nil, fmt.Errorf("error scanning user score: %v", err)
		}
		leaderboard = append(leaderboard, userScore)
	}
	return leaderboard, nil
}

func (s *PositionTableService) GetUserScoresByGroup(userID int, groupID int) ([]GroupUserScore, error) {
	query := `
		SELECT u.user_id, u.first_name, u.last_name, ug.group_id, SUM(s.points) AS points
		FROM User u
		JOIN Scores s ON u.user_id = s.user_id
		JOIN User_UserGroups ug ON u.user_id = ug.user_id
		WHERE u.user_id = ? AND ug.group_id = ?
		GROUP BY u.user_id, u.first_name, u.last_name, ug.group_id
		ORDER BY points DESC
	`
	rows, err := database.QueryRowsDBParams(query, userID, groupID)
	if err != nil {
		utils.ErrorLogger.Println("Error querying user scores by group:", err)
		return nil, fmt.Errorf("error querying user scores by group: %v", err)
	}
	defer rows.Close()

	var userScores []GroupUserScore
	for rows.Next() {
		var groupUserScore GroupUserScore
		if err := rows.Scan(&groupUserScore.UserID, &groupUserScore.FirstName, &groupUserScore.LastName, &groupUserScore.GroupID, &groupUserScore.Points); err != nil {
			utils.ErrorLogger.Println("Error scanning user score:", err)
			return nil, fmt.Errorf("error scanning user score: %v", err)
		}
		userScores = append(userScores, groupUserScore)
	}
	return userScores, nil
}
