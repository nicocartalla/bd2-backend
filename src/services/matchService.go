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

func (r *MatchService) GetAllMatchesByChampionshipID(championshipID int) ([]models.Match, error) {
    query := "SELECT * FROM GameMatch WHERE championship_id = ? ORDER BY match_date ASC"
    rows, err := database.QueryRowsDBParams(query, championshipID)
    if err != nil {
        utils.ErrorLogger.Println("Error getting results: ", err)
        return nil, fmt.Errorf("error getting results: %v", err)
    }
    defer rows.Close()

    var results []models.Match
    for rows.Next() {
        var result models.Match
        err = rows.Scan(&result.MatchID, &result.MatchDate, &result.TeamLocalID, &result.TeamVisitorID, &result.GoalsLocal, &result.GoalsVisitor, &result.ChampionshipID, &result.StageID, &result.GroupSID)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning result: ", err)
            return nil, fmt.Errorf("error scanning result: %v", err)
        }
        results = append(results, result)
    }
    return results, nil
}

func (r *MatchService) GetAllPlayedMatchesByChampionshipID(championshipID int) ([]models.Match, error) {
    query := "SELECT * FROM GameMatch WHERE championship_id = ? AND goals_local IS NOT NULL AND goals_visitor IS NOT NULL ORDER BY match_date ASC"
    rows, err := database.QueryRowsDBParams(query, championshipID)
    if err != nil {
        utils.ErrorLogger.Println("Error getting results: ", err)
        return nil, fmt.Errorf("error getting results: %v", err)
    }
    defer rows.Close()

    var results []models.Match
    for rows.Next() {
        var result models.Match
        err = rows.Scan(&result.MatchID, &result.MatchDate, &result.TeamLocalID, &result.TeamVisitorID, &result.GoalsLocal, &result.GoalsVisitor, &result.ChampionshipID, &result.StageID, &result.GroupSID)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning result: ", err)
            return nil, fmt.Errorf("error scanning result: %v", err)
        }
        results = append(results, result)
    }
    return results, nil
}


func (r *MatchService) GetNotPlayedMatchesByChampionshipID(championshipID int) ([]models.Match, error) {
    query := "SELECT * FROM GameMatch WHERE championship_id = ? AND match_date > NOW() + INTERVAL (SELECT hours_until_match FROM Utils) HOUR ORDER BY match_date ASC"
    rows, err := database.QueryRowsDBParams(query, championshipID)
    if err != nil {
        utils.ErrorLogger.Println("Error getting matches to play: ", err)
        return nil, fmt.Errorf("error getting matches to play: %v", err)
    }
    defer rows.Close()

    var matches []models.Match
    for rows.Next() {
        var match models.Match
        err = rows.Scan(&match.MatchID, &match.MatchDate, &match.TeamLocalID, &match.TeamVisitorID, &match.GoalsLocal, &match.GoalsVisitor, &match.ChampionshipID, &match.StageID, &match.GroupSID)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning match: ", err)
            return nil, fmt.Errorf("error scanning match: %v", err)
        }
        matches = append(matches, match)
    }
    return matches, nil
}


func (r *MatchService) GetMatchResult(matchID int) (models.Match, error) {
    var result models.Match

    query := "SELECT * FROM GameMatch WHERE match_id = ?"
    row, err := database.QueryRowDB(query, matchID)
    if err != nil {
        utils.ErrorLogger.Println("Error getting result: ", err)
        return models.Match{}, fmt.Errorf("error getting result: %v", err)
    }
    err = row.Scan(&result.MatchID, &result.MatchDate, &result.TeamLocalID, &result.TeamVisitorID, &result.GoalsLocal, &result.GoalsVisitor, &result.ChampionshipID, &result.StageID, &result.GroupSID)
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
    query := "INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, championship_id, stage_id, group_s_id) VALUES ( ?, ?, ?, ?, ?, ?)"
    // Print all parameters 
    utils.InfoLogger.Println("Inserting match with parameters: ", match.MatchDate, match.TeamLocalID, match.TeamVisitorID, match.ChampionshipID, match.StageID, match.GroupSID)
    result, err := database.InsertDBParams(query, match.MatchDate, match.TeamLocalID, match.TeamVisitorID, match.ChampionshipID, match.StageID, match.GroupSID)
    if err != nil {
        utils.ErrorLogger.Println("Error inserting match: ", err)
        return 0, fmt.Errorf("error inserting match: %v", err)
    }
    return result, nil
}

func (r *MatchService) UpdateMatch(match models.Match) (int64, error) {
    query := `
        UPDATE GameMatch
        SET match_date = ?, team_local_id = ?, team_visitor_id = ?, championship_id = ?, stage_id = ?, group_s_id = ?, goals_local = ?, goals_visitor = ?
        WHERE match_id = ?`

    result , err := database.UpdateDBParams(query, match.MatchDate, match.TeamLocalID, match.TeamVisitorID, match.ChampionshipID, match.StageID, match.GroupSID, match.GoalsLocal, match.GoalsVisitor, match.MatchID)
    if err != nil {
        utils.ErrorLogger.Println("Error updating match: ", err)
        return 0, fmt.Errorf("error updating match: %d , %v", result, err)
    }
    return int64(match.MatchID), nil
}

func (r *MatchService) DeleteMatch(matchID int) (int64, error) {
    query := fmt.Sprintf("DELETE FROM GameMatch WHERE match_id = %d", matchID)
    id, err := database.InsertDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error deleting match: ", err)
        return 0, fmt.Errorf("error deleting match: %v , %d" , err, id)
    }
    return int64(matchID), nil
}

func (r *MatchService) InsertResult(match models.Match) (int, error) {
    query := fmt.Sprintf("UPDATE GameMatch SET goals_local = %d, goals_visitor = %d WHERE match_id = %d", *match.GoalsLocal, *match.GoalsVisitor, match.MatchID)
    rowsAffected, err := database.UpdateDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error updating result: ", err)
        return 0, fmt.Errorf("error updating result: %v", err)
    }
    if rowsAffected == 0 {
        checkQuery := "SELECT COUNT(1) FROM GameMatch WHERE match_id = ?"
        count, err := database.QueryRowDB(checkQuery, match.MatchID)
        if err != nil {
            utils.ErrorLogger.Println("Error checking match existence: ", err)
            return 0, fmt.Errorf("error checking match existence: %v", err)
        }
        var matchExists int
        err = count.Scan(&matchExists)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning match existence: ", err)
            return 0, fmt.Errorf("error scanning match existence: %v", err)
        }

        if matchExists == 0 {
            return 0, fmt.Errorf("no match found with ID: %d", match.MatchID)
        } else {
            return 0, fmt.Errorf("no changes made to match with ID: %d, results might be the same as before", match.MatchID)
        }
    }
    return match.MatchID, nil
}



// Get match not played yet, return true if match exists and has not been played yet
func (r *MatchService) IsMatchUpcoming(matchID int) (bool, error) {
    var match models.Match
    query := "SELECT match_id FROM GameMatch WHERE match_id = ? AND match_date > NOW() + INTERVAL (SELECT hours_until_match FROM Utils) HOUR"
    row, err := database.QueryRowDB(query, matchID)
    if err != nil {
        utils.ErrorLogger.Println("Error getting match to play: ", err)
        return false, fmt.Errorf("error getting match to play: %v", err)
    }
    err = row.Scan(&match.MatchID)
    if err == sql.ErrNoRows {
        return false, fmt.Errorf("no match found with ID: %d", matchID)
    } else if err != nil {
        utils.ErrorLogger.Println("Error scanning match: ", err)
        return false, fmt.Errorf("error scanning match: %v", err)
    }
    return true, nil
}


