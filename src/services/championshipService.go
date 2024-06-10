package services

import (
    "bd2-backend/src/database"
    "bd2-backend/src/models"
    "bd2-backend/src/utils"
    "database/sql"
    "fmt"
)

type ChampionshipService struct {
    Championship models.Championship
}

//Validate Championship - Check if championship exists MOVER A CHAMPIONSHIP SERVICE CUANDO SE CREE
func (r *ChampionshipService) ValidateChampionship(championshipID int) (bool, error) {
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


func (r *ChampionshipService) GetAllChampionships() ([]models.Championship, error) {
    query := "SELECT * FROM Championships"
    rows, err := database.QueryDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error getting championships: ", err)
        return nil, fmt.Errorf("error getting championships: %v", err)
    }
    defer rows.Close()

    var championships []models.Championship
    for rows.Next() {
        var championship models.Championship
        err = rows.Scan(&championship.ChampionshipID, &championship.Name, &championship.Year, &championship.Country, &championship.ChampionshipType, &championship.Champion, &championship.Subchampion)
        if err != nil {
            utils.ErrorLogger.Println("Error scanning championship: ", err)
            return nil, fmt.Errorf("error scanning championship: %v", err)
        }
        championships = append(championships, championship)
    }
    return championships, nil
}

func (r *ChampionshipService) GetChampionshipByID(id int) (models.Championship, error) {
    var championship models.Championship
    query := "SELECT * FROM Championships WHERE championship_id = ?"
    row, err := database.QueryRowDB(query, id)
    if err != nil {
        return championship, fmt.Errorf("error getting championship: %v", err)
    }
    err = row.Scan(&championship.ChampionshipID, &championship.Name, &championship.Year, &championship.Country, &championship.ChampionshipType, &championship.Champion, &championship.Subchampion)
    if err == sql.ErrNoRows {
        return championship, fmt.Errorf("championship with id %d not found", id)
    } else if err != nil {
        return championship, fmt.Errorf("error scanning championship: %v", err)
    }
    return championship, nil
}

func (r *ChampionshipService) CreateChampionship(championship models.Championship) (int64, error) {
    query := "INSERT INTO Championships (name, year, country, championship_type) VALUES (?, ?, ?, ?)"
    result, err := database.InsertDBParams(query, championship.Name, championship.Year, championship.Country, championship.ChampionshipType)
    if err != nil {
        utils.ErrorLogger.Println("Error creating championship: ", err)
        return 0, fmt.Errorf("error creating championship: %v", err)
    }
    return result, nil
}

func (r *ChampionshipService) UpdateChampionship(id int, championship models.Championship) (int64, error) {
    query := "UPDATE Championships SET name = ?, year = ?, country = ?, championship_type = ?, champion = ?, subchampion = ? WHERE championship_id = ?"
    result, err := database.InsertDBParams(query, championship.Name, championship.Year, championship.Country, championship.ChampionshipType, championship.Champion, championship.Subchampion, id)
    if err != nil {
        utils.ErrorLogger.Println("Error updating championship: ", err)
        return 0, fmt.Errorf("error updating championship: %v", err)
    }
    return result, nil
}

//funcion para updetear el championship y el subchampion
func (r *ChampionshipService) UpdateChampionshipChampion(id int, champion int, subchampion int) (int64, error) {
    query := "UPDATE Championships SET champion = ?, subchampion = ? WHERE championship_id = ?"
    result, err := database.InsertDBParams(query, champion, subchampion, id)
    if err != nil {
        utils.ErrorLogger.Println("Error updating championship: ", err)
        return 0, fmt.Errorf("error updating championship: %v", err)
    }
    return result, nil
}

func (r *ChampionshipService) DeleteChampionship(id int) (int64, error) {
    query := "DELETE FROM Championships WHERE championship_id = ?"
    result, err := database.InsertDBParams(query, id)
    if err != nil {
        utils.ErrorLogger.Println("Error deleting championship: ", err)
        return 0, fmt.Errorf("error deleting championship: %v", err)
    }
    return result, nil
}
