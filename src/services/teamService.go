package services

import (
    "fmt"
	"bd2-backend/src/database"
    "bd2-backend/src/models"
    "bd2-backend/src/utils"
)

type TeamService struct {
    Team models.Team
}

func (t *TeamService) CheckTeamExistsByName(teamName string) bool {
	query := fmt.Sprintf("SELECT name as nameDB FROM Team WHERE name = '%s'", teamName)
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error checking if team exists: ", err)
		return false
	}
	defer rows.Close()

	var nameDB string
	if rows.Next() {
		err = rows.Scan(&nameDB)
		if err != nil {
			utils.ErrorLogger.Println("Error scanning team name: ", err)
			return false
		}
	}
	return teamName == nameDB
}

//Validate Team - Check if team exists
func (r *TeamService) CheckTeamExistsByID(teamID int) (bool, error) {
	var teamIDDB int
    query := ("SELECT team_id FROM Teams WHERE team_id = ?")
    row, err := database.QueryRowDB(query, teamID)
	row.Scan(&teamIDDB)
	if err != nil {
		utils.ErrorLogger.Println("Error checking if team exists: ", err)
		return false, fmt.Errorf("error checking if team exists: %v", err)
	}
    return teamID == teamIDDB, nil
}

func (t *TeamService) GetTeams() ([]models.Team, error) {
	query := "SELECT team_id, name FROM Teams"
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error getting teams: ", err)
		return nil, fmt.Errorf("error getting teams: %v", err)
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err = rows.Scan(&team.ID, &team.Name)
		if err != nil {
			utils.ErrorLogger.Println("Error scanning team: ", err)
			return nil, fmt.Errorf("error scanning team: %v", err)
		}
		teams = append(teams, team)
	}
	return teams, nil
}
