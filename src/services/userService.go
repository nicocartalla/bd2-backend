package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/hashing"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"fmt"
)

type UserService struct {
	User models.User
}

// CheckUserExistsByDocumentID checks if a user exists by document_id
func (u *UserService) CheckUserExistsByDocumentID(document_id string) bool {
	query := fmt.Sprintf("SELECT document_id FROM User WHERE document_id = '%s'", document_id)
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error checking if user exists: ", err)
	}
	i := 0
	var documentID string
	for rows.Next() {
		rows.Scan(&documentID)
		i++
	}
	return i != 0 && document_id == documentID
}

func (u *UserService) checkUserExists() bool {
	query := fmt.Sprintf("SELECT email FROM User WHERE email = '%s'", u.User.Email)
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error checking if user exists: ", err)
	}
	i := 0
	var email string
	for rows.Next() {
		rows.Scan(&email)
		i++
	}
	return i != 0 && u.User.Email == email
}

func (u *UserService) ValidateLogin() (bool, error) {
	if !u.checkUserExists() {
		return false, fmt.Errorf("user does not exist")
	}
	query := fmt.Sprintf("SELECT email, last_name, first_name, major, role_id, password FROM User WHERE email = '%s'", u.User.Email)
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println(err.Error())
	}
	i := 0
	var hashFromBD string
	for rows.Next() {
		i++
		err = rows.Scan(&u.User.Email, &u.User.LastName, &u.User.FirstName, &u.User.Major, &u.User.RoleID, &hashFromBD)
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
		}
	}
	if i == 0 {
		utils.WarningLogger.Println("Error validate login: ", u.User.Email)
		return false, fmt.Errorf("password error")
	}
	if i == 1 {
		if u.validateHash(hashFromBD) {
			return true, nil
		} else {
			utils.WarningLogger.Println("Error validate login: ", u.User.Email)
			return false, fmt.Errorf("password error")
		}
	}
	utils.ErrorLogger.Println("Multiple users with email: ", u.User.Email)
	return false, fmt.Errorf("multiple users with email: %s", u.User.Email)
}

// func (u *UserService) createAvatar() string {
//     return fmt.Sprintf("https://ui-avatars.com/api/?name=%s+%s?length=2", u.User.Email, u.User.LastName)
// }

func (u *UserService) CreateUser() (string, error) {
	if u.checkUserExists() {
		return "", fmt.Errorf("user already exists")
	}
	pswHashed, errHash := hashing.HashPassword(u.User.Password)
	if errHash != nil {
		return "", fmt.Errorf("error hashing password")
	}
	u.User.Password = pswHashed
	query := fmt.Sprintf("INSERT INTO User (document_id, email, first_name, last_name, major, password, role_id) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%d')", u.User.DocumentID, u.User.Email, u.User.FirstName, u.User.LastName, u.User.Major, u.User.Password, u.User.RoleID)
	_, err := database.InsertDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error creating user: ", err)
		return "", fmt.Errorf("error creating user: %v", err)
	}

	return u.User.DocumentID, nil
}

func (u *UserService) GetUser() (models.User, error) {
	query := fmt.Sprintf("SELECT document_id as Id, first_name as firstName, last_name as lastName, email, role_id FROM User WHERE id = %s", u.User.DocumentID)
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error getting user: ", err)
		return models.User{}, fmt.Errorf("error getting user: %v", err)
	}
	for rows.Next() {
		err = rows.Scan(&u.User.DocumentID, &u.User.FirstName, &u.User.LastName, &u.User.Email, &u.User.RoleID)
		if err != nil {
			utils.ErrorLogger.Println("Error getting user: ", err)
			return models.User{}, fmt.Errorf("error getting user: %v", err)
		}
	}
	return u.User, nil
}

func (u *UserService) validateHash(hashFromBD string) bool {
	return hashing.CheckPasswordHash(u.User.Password, hashFromBD)
}

// Add user to championship.
func (u *UserService) AddUserToChampionship(document_id string, championship_id int) error {
	query := fmt.Sprintf("INSERT INTO UserChampionship (document_id, championship_id) VALUES ('%s', %d)", document_id, championship_id)
	_, err := database.InsertDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error adding user to championship: ", err)
		return fmt.Errorf("error adding user to championship: %v", err)
	}
	return nil
}
