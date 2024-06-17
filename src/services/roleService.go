package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
)

type RoleService struct{}

func (s *RoleService) GetRoleId(roleName string) (models.Role, error) {
	var roleData models.Role
	query := "SELECT role_id FROM UserRoles where role_name = ? LIMIT 1"
	row, err := database.QueryRowDB(query, roleName)
	row.Scan(&roleData.ID)
	if err != nil {
		utils.ErrorLogger.Println("Error getting role id from UserRoles: ", err)
		return roleData, fmt.Errorf("error getting role from UserRoles: %v", err)
	}
	return roleData, nil
}