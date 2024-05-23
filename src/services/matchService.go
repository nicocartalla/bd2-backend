package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
    "database/sql"
)

type MatchService struct {
    Match models.Match
}

//Validate Championship - Check if championship exists MOVER A CHAMPIONSHIP SERVICE CUANDO SE CREE
func (r *MatchService) ValidateChampionship(championshipID int) (bool, error) {
    query := ("SELECT championship_id FROM Championship WHERE championship_id = %d")
    rows, err := database.QueryRowDB(query, championshipID)
    if err != nil || rows.Err() != nil{
        utils.ErrorLogger.Println("Error checking if championship exists: ", err)
        return false, fmt.Errorf("error checking if championship exists: %v", err)
    }
    var championshipIDDB int
    return championshipID == championshipIDDB, nil
}


func (r *MatchService) CheckAllResults() ([]models.Match, error) {
   // var matchTime time.Time
    query := "SELECT * FROM GameMatch"
    rows, err := database.QueryDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error getting results: ", err)
        return nil, fmt.Errorf("error getting results: %v", err)
    }
    defer rows.Close()

    var results []models.Match
    for rows.Next() {
        var result models.Match
        err = rows.Scan(&result.MatchID, &result.MatchDate, &result.TeamLocalID, &result.TeamVisitorID, &result.GoalsLocal, &result.GoalsVisitor, &result.ChampionshipID)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning result: ", err)
            return nil, fmt.Errorf("error scanning result: %v", err)
        }
        results = append(results, result)
    }
    return results, nil
}

func (r *MatchService) GetMatchResult(matchID int) (models.Match, error) {
    var result models.Match

    query := "SELECT match_id, match_date, team_local_id, team_local_id, goals_local, goals_visitor, championship_id FROM GameMatch WHERE match_id = ?"
    row, err := database.QueryRowDB(query, matchID)
    if err != nil {
        utils.ErrorLogger.Println("Error getting result: ", err)
        return models.Match{}, fmt.Errorf("error getting result: %v", err)
    }
    err = row.Scan(&result.MatchID, &result.MatchDate, &result.TeamLocalID, &result.TeamVisitorID, &result.GoalsLocal, &result.GoalsVisitor, &result.ChampionshipID)
    if err == sql.ErrNoRows {
		return models.Match{}, fmt.Errorf("no match found for match_id: %d", matchID)
	} else if err != nil {
		utils.ErrorLogger.Println("Error scanning result: ", err)
		return models.Match{}, fmt.Errorf("error scanning result: %v", err)
	}
    if result.GoalsLocal == nil && result.GoalsVisitor == nil {
        return models.Match{}, fmt.Errorf("match with match_id: %d has not been played yet", matchID)   
    }
    return result, nil
}

func (r *MatchService) InsertMatch(match models.Match) (int64, error) {
    query := fmt.Sprintf("INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, championship_id) VALUES ('%s', %d, %d, %d)", match.MatchDate, match.TeamLocalID, match.TeamVisitorID, match.ChampionshipID)
    id, err := database.InsertDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error inserting match: ", err)
        return 0, fmt.Errorf("error inserting match: %v", err)
    }
    return id, nil
}

func (r *MatchService) UpdateMatch(match models.Match) (int64, error) {
    query := fmt.Sprintf("UPDATE GameMatch SET match_date = '%s', team_local_id = %d, team_visitor_id = %d, championship_id = %d WHERE match_id = %d", match.MatchDate, match.TeamLocalID, match.TeamVisitorID, match.ChampionshipID, match.MatchID)
    id, err := database.InsertDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error updating match: ", err)
        return 0, fmt.Errorf("error updating match: %v", err)
    }
    return id, nil
}

func (r *MatchService) DeleteMatch(matchID int) (int64, error) {
    query := fmt.Sprintf("DELETE FROM GameMatch WHERE match_id = %d", matchID)
    id, err := database.InsertDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error deleting match: ", err)
        return 0, fmt.Errorf("error deleting match: %v", err)
    }
    return id, nil
}

func (r *MatchService) InsertResult(match models.Match) (int64, error) {
    query := fmt.Sprintf("INSERT INTO Results (match_id, goals_local, goals_visitor) VALUES (%d, %d, %d)", match.MatchID, match.GoalsLocal, match.GoalsVisitor)
    id, err := database.InsertDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error inserting result: ", err)
        return 0, fmt.Errorf("error inserting result: %v", err)
    }
    return id, nil
}