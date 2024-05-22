package services

import (
    "fmt"
	"bd2-backend/src/database"
    "bd2-backend/src/models"
    "bd2-backend/src/utils"
)

type ResultService struct {
    Result models.Result
}

// CheckALLResults
func (r *ResultService) CheckAllResults() ([]models.Result, error) {
    query := "SELECT * FROM Results"
    rows, err := database.QueryDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error getting results: ", err)
        return nil, fmt.Errorf("error getting results: %v", err)
    }
    defer rows.Close()

    var results []models.Result
    for rows.Next() {
        var result models.Result
        err = rows.Scan(&result.ResultID, &result.MatchID, &result.GoalsLocal, &result.GoalsVisitor)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning result: ", err)
            return nil, fmt.Errorf("error scanning result: %v", err)
        }
        results = append(results, result)
    }
    return results, nil
}
