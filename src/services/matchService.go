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
    var championshipIDDB int
    query := ("SELECT championship_id FROM Championships WHERE championship_id = ?")
    row, err := database.QueryRowDB(query, championshipID)
    row.Scan(&championshipIDDB)
    if err == sql.ErrNoRows {
        return false, fmt.Errorf("championship with championship_id: %d does not exist", championshipID)
    } else if err != nil {
        utils.ErrorLogger.Println("Error scanning championship: ", err)
        return false, fmt.Errorf("error scanning championship: %v", err)
    }
    return championshipID == championshipIDDB, nil
}

func (r *MatchService) ValidateMatch(matchID int) (bool, error) {
    var matchIDDB int
    query := ("SELECT match_id FROM GameMatch WHERE match_id = ?")
    row, err := database.QueryRowDB(query, matchID)
    row.Scan(&matchIDDB)
    if err == sql.ErrNoRows {
        return false, fmt.Errorf("match with match_id: %d does not exist", matchID)
    } else if err != nil {
        utils.ErrorLogger.Println("Error scanning match: ", err)
        return false, fmt.Errorf("error scanning match: %v", err)
    }
    return matchID == matchIDDB, nil
}

func (r *MatchService) GetAllMatchResults() ([]models.Match, error) {
    query := "SELECT * FROM GameMatch WHERE goals_local IS NOT NULL AND goals_visitor IS NOT NULL"
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
    query := "INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, championship_id) VALUES ( ?, ?, ?, ? )"
    result, err := database.InsertDBParams(query, match.MatchDate, match.TeamLocalID, match.TeamVisitorID, match.ChampionshipID)
    if err != nil {
        utils.ErrorLogger.Println("Error inserting match: ", err)
        return 0, fmt.Errorf("error inserting match: %v", err)
    }
    return result, nil
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


func (r *MatchService) GetMatchesNotPlayedYet() ([]models.Match, error) {
    query := "SELECT * FROM GameMatch WHERE match_date > NOW() + INTERVAL (SELECT hours_until_match FROM Utils) HOUR"
    rows, err := database.QueryDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error getting matches to play: ", err)
        return nil, fmt.Errorf("error getting matches to play: %v", err)
    }
    defer rows.Close()

    var matches []models.Match
    for rows.Next() {
        var match models.Match
        err = rows.Scan(&match.MatchID, &match.MatchDate, &match.TeamLocalID, &match.TeamVisitorID, &match.GoalsLocal, &match.GoalsVisitor, &match.ChampionshipID)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning match: ", err)
            return nil, fmt.Errorf("error scanning match: %v", err)
        }
        matches = append(matches, match)
    }
    return matches, nil
}

