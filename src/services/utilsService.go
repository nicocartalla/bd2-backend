package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
)

type UtilsService struct{}

func (s *UtilsService) GetPoints() (models.Utils, error) {
	var utilsData models.Utils
	query := "SELECT exact_match_points, correct_result_match_points, champion_points, sub_champion_points FROM Utils LIMIT 1"
	row, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error querying database: ", err)
		return utilsData, fmt.Errorf("error querying database: %v", err)
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&utilsData.ExactMatchPoints, &utilsData.CorrectResultMatchPoints, &utilsData.ChampionPoints, &utilsData.SubChampionPoints)
		if err != nil {
			utils.ErrorLogger.Println("Error scanning row: ", err)
			return utilsData, fmt.Errorf("error scanning row: %v", err)
		}
	} else {
		utils.ErrorLogger.Println("No rows found")
		return utilsData, fmt.Errorf("no rows found")
	}

	utils.InfoLogger.Println("Got points from Utils: ", utilsData)
	return utilsData, nil
}

func (s *UtilsService) GetPointsExactResult() (int, error) {
	utils.InfoLogger.Println("Getting exact match points")
	utilsData, err := s.GetPoints()
	utils.InfoLogger.Println("Got exact match points: ", utilsData)
	if err != nil {
		return 0, err
	}
	utils.InfoLogger.Println("Exact match points: ", utilsData.ExactMatchPoints)
	return utilsData.ExactMatchPoints, nil
}

func (s *UtilsService) GetPointsCorrectResult() (int, error) {
	utilsData, err := s.GetPoints()
	if err != nil {
		return 0, err
	}
	return utilsData.CorrectResultMatchPoints, nil
}

func (s *UtilsService) GetPointsChampion() (int, error) {
	utilsData, err := s.GetPoints()
	if err != nil {
		return 0, err
	}
	return utilsData.ChampionPoints, nil
}

func (s *UtilsService) GetPointsSubChampion() (int, error) {
	utilsData, err := s.GetPoints()
	if err != nil {
		return 0, err
	}
	return utilsData.SubChampionPoints, nil
}

func (s *UtilsService) UpdatePoints(exactMatchPoints int, correctResultMatchPoints int) error {
	query := "UPDATE Utils SET exact_match_points = ?, correct_result_match_points = ?"
	_, err := database.InsertDBParams(query, exactMatchPoints, correctResultMatchPoints)
	if err != nil {
		utils.ErrorLogger.Println("Error updating points in Utils: ", err)
		return fmt.Errorf("error updating points in Utils: %v", err)
	}
	return nil
}
